package service

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"time"

	config "git.druva.org/druva.com/DI-Master/config"
	"git.druva.org/druva.com/DI-Master/cqeclient"
	"git.druva.org/druva.com/DI-Master/rerror"
	"git.druva.org/druva.com/DI-Master/rtypes"
	productInfoClient "git.druva.org/druva.com/DI-Master/sdk/remote/phoenix"
	"git.druva.org/druva.com/DI-Master/sdk/spec"
	"git.druva.org/druva.com/DI-Master/svcparam/envvar"
	"git.druva.org/druva.com/DI-Master/util"
	"github.com/spf13/viper"
)

var cqeClient cqeclient.ICQEClient
var remoteClient productInfoClient.ProuductInfo
var nasURL string
var fsURL string

func init() {
	//to test api with mock data comment cqeclient.CQEClient{}
	// uncomment the apimocker.MockCQEClient{}
	//cqeClient = apimocker.MockCQEClient{}
	nasHost := viper.GetString(envvar.NASEndPoint)
	nasPort := viper.GetString(envvar.NasServerPort)
	fsHost := viper.GetString(envvar.FSEndPoint)
	fsPort := viper.GetString(envvar.FileServerPort)

	//phoenixJWTURL := viper.GetString(envvar.PhoenixJWTEndPoint)

	nasURL = nasHost + ":" + nasPort
	fsURL = fsHost + ":" + fsPort
	cqeClient = cqeclient.CQEClient{}
	//remoteClient = remoteClient.NewPhoenixClient()
	remoteClient = productInfoClient.NewPhoenixClient()

}

//input param :
//      1. customer Id (int)
//      2. Producd Id ([]int)
//      3. rtypes.RequestFilterData
//output param:
//      1. []interface{}
//for provided Customer id and product ids function collect list of total backed up data with respective data
func GetTotalBackupDataForDays(productsCustomers map[int]int, startdate int64, enddate int64, filter rtypes.RequestFilter) ([]interface{}, error) {
	var totalbackup []interface{}
	var date []string
	var cqefilter cqeclient.FilterData
	if len(productsCustomers) == 0 {
		return totalbackup, rerror.New(rerror.InvalidInputParam)
	}
	/*check for get last cube refresh date*/
	if startdate == 0 && enddate == 0 {
		lastRefreshDate, err := cqeclient.GetLastCubeRefreshDate(cqeClient)
		if err != nil {
			return totalbackup, rerror.New(rerror.InvalidCubeRefreshDate)
		} else {
			layout := "2006-01-02"
			timedata, err := time.Parse(layout, lastRefreshDate)
			if err != nil {
				return totalbackup, err
			}

			enddate = timedata.Unix()
			startdate = (enddate - 2505600) //2505600 = 24*29*60*60 time of 30 days

		}
	}

	if startdate > enddate {
		return totalbackup, rerror.Error(rerror.BadRequest)
	} else if startdate == enddate {
		date = append(date, util.GetEpochToStringdate(startdate))
	} else {
		date = util.GetEpochToStringDates(startdate, enddate)
	}
	cqefilter = ReqFilterToCqeFilter(filter)
	addAllFileExtentionsInCQEFilter(&cqefilter)
	responseData, err := cqeclient.GetTotalBackupDataForDays(cqeClient, productsCustomers, date, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetTotalBackupDataForDays", "Error:", err.Error())
		return totalbackup, err
	}
	layout := "2006-01-02" // cqe return data in string format ("yyyy-mm-dd")
	for _, data := range responseData {
		timedata, err := time.Parse(layout, data.Time)
		if err == nil && reflect.DeepEqual(cqefilter, (cqeclient.FilterData{})) {
			totalbackup = append(totalbackup, rtypes.Totalbackedupdata{Time: timedata.Unix(), TotalBackupData: int64(data.TotalBackupData)})
		} else if err == nil && false == reflect.DeepEqual(cqefilter, (cqeclient.FilterData{})) {
			totalbackup = append(totalbackup, rtypes.Totalfilterdata{Time: timedata.Unix(), FilteredBackedupData: int64(data.TotalBackupData)})
		} //else {
		// 		totalbackup = append(totalbackup, rtypes.Totalfilterdata{Time: time.Now().Unix(), FilteredBackedupData: int64(data.TotalBackupData)})
		// }
	}
	return totalbackup, nil
}

//input param :
//      1. customer Id(int)
//      2. Producd Id ([]int)
//	3. rtypes.RequestFilterData
//output param:
//      1. ResponseDataByResourceType
//
//for provided customer id and product ids function will collect the Data Resource name , id , and total backedup data
//data will be filtered if filter are provied
//
func GetDataByResourceType(productsCustomers map[int]int, epochdate int64, filter rtypes.RequestFilter) (rtypes.ResponseDataByResourceType, error) {
	var response rtypes.ResponseDataByResourceType
	var cqefilter cqeclient.FilterData
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	var date string
	date = util.GetEpochToStringdate(epochdate)
	cqefilter = ReqFilterToCqeFilter(filter)
	addAllFileExtentionsInCQEFilter(&cqefilter)
	responseData, err := cqeclient.GetDataBySourceType(cqeClient, productsCustomers, date, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataByResourceType", "Error:", err.Error())
		return response, err
	}

	if len(responseData) == 0 {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataByResourceType", "Error: get empty data from CQE")
		return response, rerror.Error(rerror.NoContent)
	}

	for _, data := range responseData {
		response.Sources = append(response.Sources, rtypes.DataByResourceType{SourceId: int64(data.SourceID), SourceName: sourcemap[data.SourceID], Totalbackupdata: int64(data.TotalBackupData)})

	}

	return response, nil
}

//input param :
//	1. Customer Id (int)
//	2. Producd Id ([]int)
//output param:
//	1. rtypes.GetDataSourceGroupsRes
//	2. error
//function will collect source groups and data source for provided poduct ids and customer id
//
func GetDataSourceGroups(productsCustomers map[int]int, epochDate int64) (rtypes.GetDataSourceGroupsRes, error) {
	if len(productsCustomers) == 0 {
		return rtypes.GetDataSourceGroupsRes{}, rerror.New(rerror.InvalidInputParam)
	}
	data, err := GetDataByResourceType(productsCustomers, epochDate, rtypes.RequestFilter{})
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDataSourceGroups", "Error:", err.Error())
		return rtypes.GetDataSourceGroupsRes{}, err
	}
	var sourceg []rtypes.SourceGroup
	sourceg = append(sourceg, rtypes.SourceGroup{SourceGroupNamen: "Server Workloads", SourceGroupId: config.ProductIdPhoenix, DataSources: data.Sources})
	return rtypes.GetDataSourceGroupsRes{DataSourceGroups: sourceg}, nil
}

//input param :
//      1. Customer Id (int)
//      2. Producd Id ([]int)
//output param:
//      1. rtypes.GetFileTypesAndExtensRes
//      2. error
//function will return file types and there extension for provided customer id
func GetFileTypesAndExtension(productsCustomers map[int]int) (rtypes.GetFileTypesAndExtensionRes, error) {
	var resp rtypes.GetFileTypesAndExtensionRes
	if len(productsCustomers) == 0 {
		return resp, rerror.New(rerror.InvalidInputParam)
	}
	dataInsertMux := &sync.Mutex{}
	var wg sync.WaitGroup
	date := "LATEST"
	var fileAndExt []rtypes.FileAndExt
	for _, fileId := range Filelist {
		var fileExtlist []string
		fileExtlist, ok := FileExts[fileId]
		if !ok {
			continue
		}
		fileType, ok := FileType[fileId]
		if !ok {
			continue
		}
		wg.Add(1)
		go func(fileType string, fileid int, fileExtlist []string) {
			defer wg.Done()
			responseData, err := cqeclient.GetExtension(cqeClient, productsCustomers, fileExtlist, date)
			if err != nil {
				logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetFileTypesAndExtens", "Error:", err.Error())
				return
			}
			if len(responseData) == 0 {
				return
			}
			var exts []rtypes.Extension
			var id int
			for _, element := range responseData {
				id = ExtId[element.ExtName]
				exts = append(exts, rtypes.Extension{ExtensionId: id, ExtensionName: element.ExtName,
					TotalBackedupData: int64(element.TotalBackupData)})
			}
			dataInsertMux.Lock()
			fileAndExt = append(fileAndExt, rtypes.FileAndExt{FileName: fileType, FileId: fileid, Extensions: exts})
			dataInsertMux.Unlock()

		}(fileType, fileId, fileExtlist)

	}
	wg.Wait()
	resp.FileTypesAndExtensions = fileAndExt
	return resp, nil
}

//input param :
//      1. Customer Id (int)
//      2. Producd Id ([]int)
//	3. date epoch (int)
//	4. number of recored required (int)
//	5. filter cqeclient.FilterData
//output param:
//      1. rtypes.GetSourceListResponse
//      2. error
//function will collect data source list and their respective backedupdata for provided poduct ids ,customer id, filter
//

func GetSourceList(productsCustomers map[int]int, epochdate int64, limit int, filter rtypes.RequestFilter, jwtwebtoken string) (rtypes.GetSourceListResponse, error) {
	var response rtypes.GetSourceListResponse
	var cqefilter cqeclient.FilterData
	ctx := context.Background()
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	var date string
	date = util.GetEpochToStringdate(epochdate)
	cqefilter = ReqFilterToCqeFilter(filter)
	addAllFileExtentionsInCQEFilter(&cqefilter)
	cqeResponsedata, err := cqeclient.GetDevicesSize(cqeClient, productsCustomers, limit, date, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetSourceList", "Error:", err.Error())
		return response, err
	}
	if len(cqeResponsedata) < 1 {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetSourceList", "ResponseList size :", len(cqeResponsedata))
		return response, err
	}
	logger.Debug(context.Background(), "Method", "Response from GetDevicesSize", "Device info ", fmt.Sprintf("%+v", cqeResponsedata))

	errorGetPhoenixJWT := 0
	var phoenixJWTerr error
	phoenixJWT := ""
	var allNasDevices []spec.NasServerInfo
	var allFsDevices []spec.FSServerInfo
	isNAS := false
	isFS := false
	for _, data := range cqeResponsedata {
		if data.SourceType == NAS {
			isNAS = true
		}
		if data.SourceType == WLserver {
			isFS = true
		}
		if isNAS && isFS {
			break
		}
	}

	phoenixJWT, phoenixJWTerr = remoteClient.GetPhoenixJWT(ctx, viper.GetString(envvar.PhoenixJWTEndPoint), jwtwebtoken)
	if phoenixJWTerr != nil {
		logger.Debug(ctx, "Method", "GetSourceList", "Failed to get GetPhoenixJWT ", envvar.PhoenixJWTEndPoint, "error: ", phoenixJWTerr.Error(), " error count: ", errorGetPhoenixJWT)
		return response, phoenixJWTerr
	}

	nasServerMap := make(map[spec.ServerKeyMap]spec.ServerValMap)
	fsServerMap := make(map[spec.ServerKeyMap]spec.ServerValMap)
	if isNAS {
		logger.Debug(context.Background(), "Method", "in NAS sourceType...")
		nasDeviceDetails := spec.NasDeviceCaptureJobDetails{phoenixJWT, ""}
		for true {
			nasInfoResp, err := remoteClient.GetAllNasServerInfo(ctx, &nasDeviceDetails)
			logger.Debug(context.Background(), "Method", "Response from GetAllNasServerInfo", "Servers info ", fmt.Sprintf("%+v", nasInfoResp))
			if err != nil {
				logger.Debug(ctx, "Method", "Failed to get GetAllNasServerInfo", "Failed to get device info for PageToken:", nasDeviceDetails.PageToken, "error: ", err.Error())
				break
			}
			allNasDevices = append(allNasDevices, nasInfoResp.Servers...)
			nasDeviceDetails.PageToken = nasInfoResp.PageToken
			if nasInfoResp.PageToken == "" {
				logger.Debug(ctx, "Method", "GetAllNasServerInfo", "Message: ", "In last pageToekn for NAS")
				break
			}
		}
	}

	if isFS {
		logger.Debug(context.Background(), "Method", "in FS sourceType...")
		fsDeviceDetails := spec.FSDeviceCaptureJobDetails{phoenixJWT, ""}
		for true {
			fsInfoResp, fsErr := remoteClient.GetAllFSServerInfo(ctx, &fsDeviceDetails)
			logger.Debug(context.Background(), "Method", "Response from GetAllFSServerInfo", "Servers info ", fmt.Sprintf("%+v", fsInfoResp))
			if fsErr != nil {
				logger.Debug(context.Background(), "package", Pkg_name, "Method", "Failed to get GetAllFSServerInfo", "Error:", fsErr.Error())
				break
			}
			allFsDevices = append(allFsDevices, fsInfoResp.Servers...)
			fsDeviceDetails.PageToken = fsInfoResp.PageToken
			if fsInfoResp.PageToken == "" {
				logger.Debug(ctx, "Method", "GetAllFSServerInfo", "Message: ", "In last pageToekn for FS")
				break
			}
		}
	}
	/*Creatting map of NAS servers */
	for _, serverInfo := range allNasDevices {
		nasServerMap[spec.ServerKeyMap{serverInfo.OrganizationID, serverInfo.StorageInfo.CsetID}] = spec.ServerValMap{serverInfo.NasDevice.ServerName, serverInfo.BackupsetName}
	}
	/*Creatting map of FS servers */
	for _, serverInfo := range allFsDevices {
		fsServerMap[spec.ServerKeyMap{serverInfo.OrganizationID, serverInfo.StorageInfo.CsetID}] = spec.ServerValMap{serverInfo.FSDevice.ServerName, serverInfo.BackupsetName}
	}
	for _, data := range cqeResponsedata {
		serverName := ""
		backupSetName := ""
		if data.SourceType == NAS {
			if value, ok := nasServerMap[spec.ServerKeyMap{int32(data.OrganisationId), int32(data.BackupSetId)}]; ok {
				serverName = value.ServerName
				backupSetName = value.BkpsetName
			}
		}
		if data.SourceType == WLserver {
			if value, ok := fsServerMap[spec.ServerKeyMap{int32(data.OrganisationId), int32(data.BackupSetId)}]; ok {
				serverName = value.ServerName
				backupSetName = value.BkpsetName
			}
		}
		if serverName == "" || backupSetName == "" {
			//Adding extra [ ] to handle AM charts escape character issue of single [ ]
			backupSetName = "[[delete : " + fmt.Sprintf("%v", data.BackupSetId) + "]]"
			//serverName = "[delete : " + fmt.Sprintf("%v", data.ServerID) + "]"
		}
		response.DataSources = append(response.DataSources, rtypes.ServerData{ServerName: serverName, TotalBackedupData: int64(data.TotalBackupData), ServerID: int64(data.ServerID), BackupSetId: data.BackupSetId, BackupSetName: backupSetName})
	}

	//server := "server" + strconv.Itoa(i)
	//backupset := "backupset" + strconv.Itoa(i)
	// response.DataSources = append(response.DataSources, rtypes.ServerData{ServerName: server, TotalBackedupData: int64(data.TotalBackupData), ServerID: int64(data.ServerID), BackupSetId: data.BackupSetId, BackupSetName: backupset})

	return response, nil
}

//input param :
//      1. Customer Id (int)
//      2. Producd Id ([]int)
//      3. date epoch (int)
//      5. filter cqeclient.FilterData
//output param:
//      1. rtypes.GetFileTypesRes
//      2. error
//function will return list of file Extension for give file type id
//
func GetFileExtension(productsCustomers map[int]int, epochdate int64, limit int, filter rtypes.RequestFilter) (rtypes.GetFileExtensionsRes, error) {
	var response rtypes.GetFileExtensionsRes
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	var cqefilter cqeclient.FilterData
	var date string
	cqefilter = ReqFilterToCqeFilter(filter)
	date = util.GetEpochToStringdate(epochdate)
	var fileExtlist []string
	var fileTypes []int

	if len(cqefilter.FileTypeIds) == 0 {
		fileTypes = Filelist
	} else {
		fileTypes = cqefilter.FileTypeIds
	}

	for _, fileTypeId := range fileTypes {
		var fileExts []string

		if len(cqefilter.FileExtenstions) >= 1 {
			fileExts = getExtForFileID(fileTypeId, cqefilter.FileExtenstions)
		} else {
			fileExts = FileExts[fileTypeId]
		}
		for _, ext := range fileExts {
			fileExtlist = append(fileExtlist, ext)
		}
	}

	if len(fileExtlist) <= 0 {
		return response, rerror.Error(rerror.InvalidFileID)
	}
	responseData, err := cqeclient.GetExtension(cqeClient, productsCustomers, fileExtlist, date, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetFileExtens", "Error:", err.Error())
		return response, err

	}

	// sort the file extension on total data in descensing order
	sort.Slice(responseData, func(i, j int) bool {
		return responseData[i].TotalBackupData > responseData[j].TotalBackupData
	})

	var extension []rtypes.Extension
	var id int
	for index, element := range responseData {
		if index >= limit && limit != 0 {
			break
		}
		id = ExtId[element.ExtName]
		extension = append(extension, rtypes.Extension{ExtensionName: element.ExtName, TotalBackedupData: int64(element.TotalBackupData), ExtensionId: id})
	}
	if len(extension) >= 1 {
		response.Extensions = extension
	}
	return response, nil
}

//input param :
//      1. Customer Id (int)
//      2. Producd Id ([]int)
//      3. date epoch (int)
//      5. filter cqeclient.FilterData
//output param:
//      1. rtypes.GetFileTypesRes
//      2. error
//function will return file type  and their respective backedupdata for provided poduct ids ,customer id, filter
//
func GetFileTypes(productsCustomers map[int]int, epochdate int64, limit int, filter rtypes.RequestFilter) (rtypes.GetFileTypesRes, error) {
	var response rtypes.GetFileTypesRes
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	dataInsertMux := &sync.Mutex{}
	var cqefilter cqeclient.FilterData
	var wg sync.WaitGroup
	var date string
	var fileInfo []rtypes.File
	var getFileTypesRes []rtypes.File
	cqefilter = ReqFilterToCqeFilter(filter)
	date = util.GetEpochToStringdate(epochdate)
	var fileTypelist []int
	if len(cqefilter.FileTypeIds) >= 1 {
		fileTypelist = cqefilter.FileTypeIds
	} else {
		fileTypelist = Filelist
	}
	for _, fileId := range fileTypelist {
		var fileExtlist []string
		if len(cqefilter.FileExtenstions) >= 1 {
			fileExtlist = getExtForFileID(fileId, cqefilter.FileExtenstions)
		} else {
			fileExtlist = FileExts[fileId]
			if len(fileExtlist) <= 0 {
				continue
			}
		}
		fileType := FileType[fileId]
		wg.Add(1)
		go func(fileType string, fileId int, fileExtlist []string) {
			defer wg.Done()
			responseData, err := cqeclient.GetExtension(cqeClient, productsCustomers, fileExtlist, date, cqefilter)
			if err != nil {
				logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetFileTypes", "Error:", err.Error()+":Failed to get data for file type:"+fileType)
				return
			}
			if len(responseData) == 0 {
				return
			}
			var totalData int64
			for _, element := range responseData {
				totalData = totalData + int64(element.TotalBackupData)
			}
			dataInsertMux.Lock()
			fileInfo = append(fileInfo, rtypes.File{FileName: fileType, FileID: fileId, TotalBackedupData: totalData})
			dataInsertMux.Unlock()

		}(fileType, fileId, fileExtlist)

	}
	wg.Wait()
	// sort File type on total data descensing order
	sort.Slice(fileInfo, func(i, j int) bool {
		return fileInfo[i].TotalBackedupData > fileInfo[j].TotalBackedupData
	})
	for index, element := range fileInfo {
		if index >= limit && limit != 0 {
			break
		}
		getFileTypesRes = append(getFileTypesRes, element)
	}
	response.FileTypes = getFileTypesRes
	return response, nil
}

//input param :
//      1. Customer Id (int)
//      2. Producd Id ([]int)
//      3. date epoch (int)
//	4. Size range (rtypes.Range)
//      5. filter cqeclient.FilterData
//output param:
//      1. rtypes.GetFileTypesRes
//      2. error
//function will return file type  and their respective backedupdata for provided poduct ids ,customer id, filter
//
func GetDistributionBySize(productsCustomers map[int]int, epochdate int64, sizeRangeList []rtypes.Range, filter rtypes.RequestFilter) (rtypes.GetDistributionBySizeRes, error) {
	var response rtypes.GetDistributionBySizeRes
	var cqefilter cqeclient.FilterData
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	var date string
	cqefilter = ReqFilterToCqeFilter(filter)
	addAllFileExtentionsInCQEFilter(&cqefilter)

	date = util.GetEpochToStringdate(epochdate)
	if len(sizeRangeList) <= 0 {
		sizeRangeList = rtypes.DefaultSizeRangelist
	}
	responseData, err := cqeclient.GetDataDistributionBySize(cqeClient, productsCustomers, sizeRangeList, date, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDistributionBySize", "Error:", err.Error()+"Failed to get data for size slab data")
		return response, err
	}
	var sizeRange []rtypes.SizeRange
	for _, element := range responseData {
		sizeRange = append(sizeRange, rtypes.SizeRange{StartByte: int64(element.StartByte), TotalBackedupData: int64(element.TotalBackedupData), EndByte: int64(element.EndByte), FileSizeSlab: element.SizeSlab})
	}
	if len(sizeRange) > 1 {
		response.DataSizeRange = sizeRange
	}
	return response, nil

}

//input param :
//      1. customer Id (int)
//      2. Producd Id ([]int)
//      3. startmonth epoch (int)
//      4. endmonth epoch(int)
//      6. filter rtypes.RequestFilter
//output param:
//      1. rtypes.DistributionByModifiedMonthRes
//      2. error
//function will return month  and their respective modified data size for provided poduct ids, customer id, filter
//
func DistributionByModifiedMonth(productsCustomers map[int]int, date int64, filter rtypes.RequestFilter) (rtypes.DistributionByModifiedMonthRes, error) {
	var response rtypes.DistributionByModifiedMonthRes
	var cqefilter cqeclient.FilterData
	var aggrigateData int64
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	cqefilter = ReqFilterToCqeFilter(filter)
	addAllFileExtentionsInCQEFilter(&cqefilter)
	responseData, err := cqeclient.DistributionByModifiedMonth(cqeClient, productsCustomers, date, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetDistributionBySize", "Error:", err.Error()+"Failed to get data for size slab data")
		return response, err
	}

	sort.SliceStable(responseData, func(i, j int) bool {
		return responseData[i].Month < responseData[j].Month
	})

	var sizebymonth []rtypes.SizeByMonth
	var firstMonthList []rtypes.SizeByMonth
	oldestMonth := 13
	currentYear := 0
	for _, element := range responseData {
		if element.Month != 0 {
			t := time.Unix(element.Month, 0)
			month := int(t.Month())
			year := int(t.Year())
			yearstr := strconv.Itoa(year)
			if month < oldestMonth {
				oldestMonth = month
				currentYear = year
			}
			sizebymonth = append(sizebymonth, rtypes.SizeByMonth{Month: MonthMap[month] + " " + yearstr[len(yearstr)-2:], ModifiedSize: element.ModifiedSize})
		} else {
			aggrigateData = element.ModifiedSize
		}
	}
	if len(responseData) > 0 {
		var yearstr string
		var monthAggrigate int
		if oldestMonth == 13 {
			t := time.Now()
			monthAggrigate = int(t.Month())
			year := int(t.Year())
			yearstr = strconv.Itoa(year)
		} else {
			monthAggrigate = oldestMonth - 1
			//if oldest month is jan
			if monthAggrigate == 0 {
				monthAggrigate = 12
				yearstr = strconv.Itoa(currentYear - 1)
			} else {
				yearstr = strconv.Itoa(currentYear)
			}
		}
		firstMonthList = append(firstMonthList, rtypes.SizeByMonth{Month: "< " + MonthMap[monthAggrigate] + " " + yearstr[len(yearstr)-2:], ModifiedSize: aggrigateData})
		sizebymonth = append(firstMonthList, sizebymonth...)
		//sizebymonth = append(sizebymonth, rtypes.SizeByMonth{Month: "< " + MonthMap[monthAggrigate] + " " + yearstr[len(yearstr)-2:], ModifiedSize: aggrigateData})

	}
	if len(sizebymonth) >= 1 {
		response.ModefiedSizeByMonths = sizebymonth
	}
	return response, nil
}

//input param :
//      1. Customer Id (int)
//      2. Producd Id ([]int)
//      3. date epoch (int)
//      5. filter cqeclient.FilterData
//output param:
//      1. rtypes.GetOrganizationsRes
//      2. error
//function will return Organizations  and their respective backedupdata for provided poduct ids ,customer id, filter
//
func GetOrganizations(productsCustomers map[int]int, epochdate int64, limit int, filter rtypes.RequestFilter, jwt string) (rtypes.GetOrganizationsRes, error) {
	var response rtypes.GetOrganizationsRes
	var cqefilter cqeclient.FilterData
	if len(productsCustomers) == 0 {
		return response, rerror.New(rerror.InvalidInputParam)
	}
	var date string
	date = util.GetEpochToStringdate(epochdate)
	cqefilter = ReqFilterToCqeFilter(filter)
	addAllFileExtentionsInCQEFilter(&cqefilter)

	responseData, err := cqeclient.GetOrganizationslist(cqeClient, productsCustomers, date, limit, cqefilter)
	if err != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetOrganizations", "Error:", err.Error()+"Failed to get data for size slab data")
		return response, err
	}

	//Changes for the Phoenix Orgs.
	ctx := context.Background()
	phoenixJWT, phoenixErr := remoteClient.GetPhoenixJWT(ctx, viper.GetString(envvar.PhoenixJWTEndPoint), jwt)
	if phoenixErr != nil {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetOrganizations", "Error in calling GetPhoenixJWT", phoenixErr.Error())
		return response, phoenixErr
	}
	var orgList []rtypes.OrgInfo
	pageToken := ""
	for true {
		orgs, orgError := remoteClient.GetOrganizationList(ctx, viper.GetString(envvar.PhoenixListOrgEndPoint), phoenixJWT, pageToken)
		if orgError != nil {
			logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetOrganizations", "Error in calling GetOrganizationList", orgError.Error())
			return response, orgError
		}
		orgList = append(orgList, orgs.OrgList...)
		if orgs.PageToken == "" {
			break
		}
		pageToken = orgs.PageToken
	}
	var Orgs []rtypes.Organization
	for _, element := range responseData {
		orgName := ""
		for _, org := range orgList {
			if element.OrganisationID == (org.OrganizationId) {
				orgName = org.OrganizationName
			}
		}
		if orgName == "" {
			//Not found Org
			//Adding extra [ ] to handle AM charts escape character issue of single [ ]
			orgName = "[[delete : " + fmt.Sprintf("%v", element.OrganisationID) + "]]"
		}
		Orgs = append(Orgs, rtypes.Organization{OrganizationId: element.OrganisationID, OrganizationName: orgName, TotalBackedupData: element.Size})
	}
	if len(Orgs) > 0 {
		response.Organizations = Orgs
	} else {
		logger.Debug(context.Background(), "package", Pkg_name, "Method", "GetOrganizations", "Error: get empty data from CQE")
		return response, rerror.Error(rerror.NoContent)
	}

	return response, nil

}

//function to convert request filter to Cqe filter
func ReqFilterToCqeFilter(reqFilter rtypes.RequestFilter) cqeclient.FilterData {
	var subSources []int
	var cqeFilter cqeclient.FilterData
	if len(reqFilter.SubSourceIds) > 0 {
		for _, elementid := range reqFilter.SubSourceIds {
			if len(sourcemap[elementid]) > 0 {
				subSources = append(subSources, elementid)
			}
		}
	}
	cqeFilter.SourceId = reqFilter.SourceId
	cqeFilter.SubSources = subSources
	cqeFilter.ProfileIds = reqFilter.ProfileIds
	cqeFilter.OrganisationIds = reqFilter.OrganisationIds
	cqeFilter.FileTypeIds = reqFilter.FileTypeIds
	cqeFilter.BackupSetIds = reqFilter.BackupSetIds
	cqeFilter.DeviceIds = reqFilter.DeviceIds
	if len(reqFilter.FileExtenstionIds) > 0 {
		for _, fileid := range reqFilter.FileExtenstionIds {
			extName, ok := Ext[fileid]
			if ok {
				cqeFilter.FileExtenstions = append(cqeFilter.FileExtenstions, extName)
			}
		}
	}
	return cqeFilter
}

//function will return filter extension for given file id
func getExtForFileID(fileID int, Extensions []string) []string {
	var FileAllExt []string
	var CommonExt []string
	FileAllExt = FileExts[fileID]
	for _, fileExt := range FileAllExt {
		for _, filterExt := range Extensions {
			if fileExt == filterExt {
				CommonExt = append(CommonExt, fileExt)
			}
		}
	}
	return CommonExt
}

//function to get all extenstion from the file types
func addAllFileExtentionsInCQEFilter(cqefilter *cqeclient.FilterData) {

	if len(cqefilter.FileExtenstions) == 0 {
		for _, fileTypeId := range cqefilter.FileTypeIds {
			fileExts := FileExts[fileTypeId]
			for _, ext := range fileExts {
				cqefilter.FileExtenstions = append(cqefilter.FileExtenstions, ext)
			}
		}
	}
}
