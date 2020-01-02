package service

import (
	"reflect"
	"testing"

	"git.druva.org/druva.com/DI-Master/apimocker"
	"git.druva.org/druva.com/DI-Master/config"
	"git.druva.org/druva.com/DI-Master/cqeclient"
	"git.druva.org/druva.com/DI-Master/rerror"
	"git.druva.org/druva.com/DI-Master/rtypes"
	"github.com/golang/mock/gomock"
)

func TestGetTotalBackupDataForDays(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		startdate         int64
		enddate           int64
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetTotalBackupDataForDays without filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 103},
			startdate:         0,
			enddate:           0,
			Errflag:           false,
		},
		{
			name:              "Positive | GetTotalBackupDataForDays with filter (Success)",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 103},
			startdate:         0,
			enddate:           0,
			filter:            rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{config.ProductIdInsync, config.ProductIdPhoenix}},
			Errflag:           false,
		},
		{
			name:              "Positive | GetTotalBackupDataForDays with date range (Success)",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 103},
			startdate:         1222222,
			enddate:           2333333,
			filter:            rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{config.ProductIdInsync, config.ProductIdPhoenix}},
			Errflag:           false,
		},
		{
			name:              "Positive | GetTotalBackupDataForDays blank dates (Success)",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 103},
			startdate:         0,
			enddate:           0,
			filter:            rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{config.ProductIdInsync, config.ProductIdPhoenix}},
			Errflag:           false,
		},
		{
			name:              "Negative | GetTotalBackupDataForDays invalid date (Negative)",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 103},
			startdate:         123,
			enddate:           12,
			filter:            rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{config.ProductIdInsync, config.ProductIdPhoenix}},
			Errflag:           true,
			wantError:         rerror.Error(rerror.BadRequest),
		},
		{
			name:      "Negative | GetTotalBackupDataForDays invalid customerid (Negative)",
			startdate: 123,
			enddate:   12,
			filter:    rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{config.ProductIdInsync, config.ProductIdPhoenix}},
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		// {
		// 	name:              "Negative | GetTotalBackupDataForDays invalid customerid (Negative)",
		// 	productsCustomers: map[int]int{config.ProductIdPhoenix: 100},
		// 	startdate:         0,
		// 	enddate:           0,
		// 	filter:            rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{config.ProductIdInsync, config.ProductIdPhoenix}},
		// 	Errflag:           true,
		// 	wantError:         rerror.New(rerror.InvalidCubeRefreshDate),
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetTotalBackupDataForDays(tt.productsCustomers, tt.startdate, tt.enddate, tt.filter)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}
}

func TestGetDataByResourceType(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		epochdate         int64
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetDataByResourceType without filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 201},
			epochdate:         0,
			Errflag:           false,
		},
		{
			name:              "Positive | GetDataByResourceType with filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 201},
			epochdate:         0,
			filter:            rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{1, 2}},
			Errflag:           false,
		},
		{
			name: "Negative | GetDataByResourceType empty customerID (Negative)",
			//productsCustomers: map[int]int{config.ProductIdPhoenix: 0},
			epochdate: 0,
			filter:    rtypes.RequestFilter{SourceId: 12289, SubSourceIds: []int{1, 2}},
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		{
			name:              "Negative | GetDataByResourceType with invalid id 100",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 100},
			epochdate:         0,
			Errflag:           true,
			wantError:         rerror.Error(rerror.NoContent),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetDataByResourceType(tt.productsCustomers, tt.epochdate, tt.filter)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}
}

func TestGetDataSourceGroups(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetDataSourceGroups Success Case",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 201},
			Errflag:           false,
		},
		{
			name:      "Negative | Empty customer ID",
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetDataSourceGroups(tt.productsCustomers, 0)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}

}

func TestGetFileTypesAndExtens(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetFileTypesAndExtens Success case",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 305},
			Errflag:           false,
		},
		{
			name:              "Positive |GetFileTypesAndExtens Success case",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 431},
			Errflag:           false,
		},
		{
			name:      "Negative |GetFileTypesAndExtens empty customer ID",
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		{
			name:      "Negative |GetFileTypesAndExtens empty product ID",
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetFileTypesAndExtension(tt.productsCustomers)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}

}

func TestGetSourceList(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	remoteClient = apimocker.NewMockPhoenixClient()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		SourceId          int
		epochdate         int64
		limit             int
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetSourceList test 1(windows/Linux)",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 302},
			SourceId:          2,
			epochdate:         0,
			limit:             10,
			Errflag:           false,
		},
		{
			name:              "Positive | GetSourceList test 2(NAS)",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 302},
			SourceId:          1,
			epochdate:         0,
			limit:             10,
			Errflag:           false,
		},
		{
			name:              "Negative | GetSourceList test invalid Id",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 100},
			SourceId:          1,
			epochdate:         0,
			limit:             10,
			Errflag:           true,
			wantError:         rerror.Error(rerror.InternalError),
		},
		{
			name:      "Negative | GetSourceList test empty id",
			SourceId:  1,
			epochdate: 0,
			limit:     10,
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetSourceList(tt.productsCustomers, tt.epochdate, tt.limit, tt.filter, "")
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}
}

func TestGetFileExtens(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		FileTypeId        int
		limit             int
		epochdate         int64
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetFileExtens Success case",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 305},
			FileTypeId:        1,
			epochdate:         0,
			limit:             5,
			Errflag:           false,
		},
		{
			name:              "Positive | GetFileExtens with filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 305},
			FileTypeId:        1,
			filter:            rtypes.RequestFilter{FileTypeIds: []int{1}, FileExtenstionIds: []int{5, 6}},
			epochdate:         0,
			Errflag:           false,
		},
		{
			name:       "Negative | GetFileExtens test invalid customer ID",
			FileTypeId: 1,
			filter:     rtypes.RequestFilter{FileTypeIds: []int{1}, FileExtenstionIds: []int{5, 6}},
			epochdate:  0,
			Errflag:    true,
			wantError:  rerror.New(rerror.InvalidInputParam),
		},
		{
			name:              "Positive | GetFileExtens with fileID",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 3050},
			FileTypeId:        1,
			epochdate:         0,
			Errflag:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetFileExtension(tt.productsCustomers, tt.epochdate, tt.limit, tt.filter)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}
}

func TestGetFileTypes(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		epochdate         int64
		limit             int
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetFileTypes Success case ",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 305},
			epochdate:         0,
			limit:             5,
			Errflag:           false,
		},
		{
			name:              "Positive | GetFileTypes with filter1",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 305},
			epochdate:         0,
			filter:            rtypes.RequestFilter{FileTypeIds: []int{1}, FileExtenstionIds: []int{1, 2}},
			Errflag:           false,
		},
		{
			name:              "Positive | GetFileTypes test with filter2",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 305},
			epochdate:         0,
			filter:            rtypes.RequestFilter{FileTypeIds: []int{1}},
			Errflag:           false,
		},
		{
			name:      "Negativev | GetFileTypes test with filter3 with invalid customer id",
			epochdate: 0,
			filter:    rtypes.RequestFilter{FileTypeIds: []int{1}, FileExtenstionIds: []int{1, 2}},
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		{
			name:      "Positive | GetFileTypes test with filter empty customer ID",
			epochdate: 305,
			filter:    rtypes.RequestFilter{FileTypeIds: []int{1234}, FileExtenstionIds: []int{1, 2}},
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		{
			name:              "Positive | GetFileTypes test with filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 3050},
			epochdate:         0,
			filter:            rtypes.RequestFilter{FileTypeIds: []int{1}},
			Errflag:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var output []GetDataDistributionBySizeRes
			var err error
			_, err = GetFileTypes(tt.productsCustomers, tt.epochdate, tt.limit, tt.filter)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}

}

func TestGetDistributionBySize(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		epochdate         int64
		filter            rtypes.RequestFilter
		sizeRangeList     []rtypes.Range
		Errflag           bool
		wantError         error
	}{
		{
			name:              "Positive | GetDistributionBySize Success Case",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 445},
			Errflag:           false,
		},
		{
			name:              "Positive | GetDistributionBySize with size range",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 445},
			sizeRangeList:     []rtypes.Range{rtypes.Range{StartByte: 1048576, EndByte: 10485760}},
			Errflag:           false,
		},
		{
			name:      "Negative | GetDistributionBySize empty customerID",
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			_, err = GetDistributionBySize(tt.productsCustomers, tt.epochdate, tt.sizeRangeList, tt.filter)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}
}

func TestDistributionByModifiedMonth(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		date              int64
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "positive | valid month epoch",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 665},
			Errflag:           false,
			date:              111111111111,
		},
		{
			name:              "positive | valid month epoch",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 965},
			date:              111111111111,
			Errflag:           false,
		},
		{
			name:      "Negative | empty customer id",
			date:      111111111111,
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		{
			name:              "Negative | with filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 665},
			date:              111111111111,
			Errflag:           false,
			filter:            rtypes.RequestFilter{FileTypeIds: []int{1234}, FileExtenstionIds: []int{1, 2}},
			wantError:         rerror.New(rerror.InvalidInputParam),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			_, err = DistributionByModifiedMonth(tt.productsCustomers, tt.date, tt.filter)
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}

}

func TestGetOrganizations(t *testing.T) {
	cqeClient = apimocker.MockCQEClient{}
	remoteClient = apimocker.NewMockPhoenixClient()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name              string
		productsCustomers map[int]int
		limit             int
		epochdate         int64
		filter            rtypes.RequestFilter
		Errflag           bool
		wantError         error
	}{
		{
			name:              "positive | valid input paramter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 765},
			limit:             5,
			epochdate:         111111,
			Errflag:           false,
		},
		{
			name:      "positive | empty customer ID",
			epochdate: 111111,
			Errflag:   true,
			wantError: rerror.New(rerror.InvalidInputParam),
		},
		{
			name:              "positive | valid input paramter with filter",
			productsCustomers: map[int]int{config.ProductIdPhoenix: 765},
			filter:            rtypes.RequestFilter{FileTypeIds: []int{1234}, FileExtenstionIds: []int{1, 2}},
			epochdate:         111111,
			Errflag:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			_, err = GetOrganizations(tt.productsCustomers, tt.epochdate, tt.limit, tt.filter, "")
			if tt.Errflag {
				if err == nil {
					t.Errorf("cqeLayer: Error Required Not Raised")
				}
				switch err.(type) {
				case rerror.RError:
					if !reflect.DeepEqual(tt.wantError, err) {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}

				case error:
					if tt.wantError.Error() != err.Error() {
						t.Errorf("cqeLayer: Wrong Error got:%v ,want:%v ", err, tt.wantError)
					}
				}

			} else {
				if err != nil {
					t.Errorf("cqeLayer: error not required Error:%v", err)
				}
			}

		})
	}

}

func TestReqFilterToCqeFilter(t *testing.T) {
	tests := []struct {
		name      string
		resfilter rtypes.RequestFilter
		cqefilter cqeclient.FilterData
	}{
		{
			name:      "ReqFilterToCqeFilter Success case",
			resfilter: rtypes.RequestFilter{SourceId: 1, SubSourceIds: []int{4}, ProfileIds: []int{1, 2, 3}, OrganisationIds: []int{1, 2, 3}, FileTypeIds: []int{1}, DeviceIds: []int{1, 2, 3}, FileExtenstionIds: []int{1, 2}},
			cqefilter: cqeclient.FilterData{SourceId: 1, SubSources: []int{4}, ProfileIds: []int{1, 2, 3}, OrganisationIds: []int{1, 2, 3}, FileTypeIds: []int{1}, DeviceIds: []int{1, 2, 3}, FileExtenstions: []string{"3dm", "3ds"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ReqFilterToCqeFilter(tt.resfilter)
			if !reflect.DeepEqual(tt.cqefilter, output) {
				t.Errorf("cqeLayer: Wrong output got:%v ,want:%v ", output, tt.cqefilter)
			}
		})
	}

}

func TestgetExtForFileID(t *testing.T) {
	tests := []struct {
		name       string
		fileid     int
		extensions []string
		output     []string
	}{
		{
			name:       "Positive | getExtForFileID Success case1",
			fileid:     1,
			extensions: []string{".3dm", ".3ds", ".pdf"},
			output:     []string{".3dm", ".3ds"},
		},
		{
			name:       "Positive | getExtForFileID Success case2",
			fileid:     1,
			extensions: []string{".pdf"},
			output:     []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := getExtForFileID(tt.fileid, tt.extensions)
			if !reflect.DeepEqual(tt.output, output) {
				t.Errorf("cqeLayer: Wrong output got:%v ,want:%v ", output, tt.output)
			}
		})
	}
}
