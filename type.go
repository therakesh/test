package service

const (
	WLserver = 4 //1
	NAS      = 9 //2
	Linux    = 3
)

const (
	STATUS200 = 200
)

var sourcemap = map[int]string{
	NAS:      "NAS",
	WLserver: "Windows/Linux Servers",
	Linux:    "Linux",
}

const (
	EXT_STR_3dm           = "3dm"
	EXT_STR_3ds           = "3ds"
	EXT_STR_max           = "max"
	EXT_STR_obj           = "obj"
	EXT_STR_7z            = "7z"
	EXT_STR_cbr           = "cbr"
	EXT_STR_deb           = "deb"
	EXT_STR_gz            = "gz"
	EXT_STR_pkg           = "pkg"
	EXT_STR_rar           = "rar"
	EXT_STR_rpm           = "rpm"
	EXT_STR_sitx          = "sitx"
	EXT_STR_tar_gz        = "tar.gz"
	EXT_STR_zip           = "zip"
	EXT_STR_zipx          = "zipx"
	EXT_STR_aif           = "aif"
	EXT_STR_iff           = "iff"
	EXT_STR_m3u           = "m3u"
	EXT_STR_m4a           = "m4a"
	EXT_STR_mid           = "mid"
	EXT_STR_mp3           = "mp3"
	EXT_STR_mpa           = "mpa"
	EXT_STR_wav           = "wav"
	EXT_STR_wma           = "wma"
	EXT_STR_bak           = "bak"
	EXT_STR_old           = "old"
	EXT_STR_tmp           = "tmp"
	EXT_STR_dwg           = "dwg"
	EXT_STR_dxf           = "dxf"
	EXT_STR_csv           = "csv"
	EXT_STR_dat           = "dat"
	EXT_STR_ged           = "ged"
	EXT_STR_key           = "key"
	EXT_STR_keychain      = "keychain"
	EXT_STR_sdf           = "sdf"
	EXT_STR_tar           = "tar"
	EXT_STR_tax2016       = "tax2016"
	EXT_STR_tax2018       = "tax2018"
	EXT_STR_vcf           = "vcf"
	EXT_STR_xml           = "xml"
	EXT_STR_accdb         = "accdb"
	EXT_STR_db            = "db"
	EXT_STR_dbf           = "dbf"
	EXT_STR_mdb           = "mdb"
	EXT_STR_pdb           = "pdb"
	EXT_STR_sql           = "sql"
	EXT_STR_sqlite_wal    = "sqlite-wal"
	EXT_STR_c             = "c"
	EXT_STR_class         = "class"
	EXT_STR_cpp           = "cpp"
	EXT_STR_cs            = "cs"
	EXT_STR_dtd           = "dtd"
	EXT_STR_fla           = "fla"
	EXT_STR_h             = "h"
	EXT_STR_java          = "java"
	EXT_STR_lua           = "lua"
	EXT_STR_m             = "m"
	EXT_STR_mpp           = "mpp"
	EXT_STR_pl            = "pl"
	EXT_STR_py            = "py"
	EXT_STR_sh            = "sh"
	EXT_STR_sln           = "sln"
	EXT_STR_swift         = "swift"
	EXT_STR_vb            = "vb"
	EXT_STR_vcxproj       = "vcxproj"
	EXT_STR_xcodeproj     = "xcodeproj"
	EXT_STR_bin           = "bin"
	EXT_STR_cue           = "cue"
	EXT_STR_dmg           = "dmg"
	EXT_STR_iso           = "iso"
	EXT_STR_mdf           = "mdf"
	EXT_STR_toast         = "toast"
	EXT_STR_vcd           = "vcd"
	EXT_STR_doc           = "doc"
	EXT_STR_docx          = "docx"
	EXT_STR_odt           = "odt"
	EXT_STR_pages         = "pages"
	EXT_STR_rtf           = "rtf"
	EXT_STR_tex           = "tex"
	EXT_STR_txt           = "txt"
	EXT_STR_vsd           = "vsd"
	EXT_STR_wpd           = "wpd"
	EXT_STR_wps           = "wps"
	EXT_STR_hqx           = "hqx"
	EXT_STR_mim           = "mim"
	EXT_STR_uue           = "uue"
	EXT_STR_apk           = "apk"
	EXT_STR_app           = "app"
	EXT_STR_bat           = "bat"
	EXT_STR_cgi           = "cgi"
	EXT_STR_com           = "com"
	EXT_STR_exe           = "exe"
	EXT_STR_gadget        = "gadget"
	EXT_STR_jar           = "jar"
	EXT_STR_wsf           = "wsf"
	EXT_STR_fnt           = "fnt"
	EXT_STR_fon           = "fon"
	EXT_STR_otf           = "otf"
	EXT_STR_ttf           = "ttf"
	EXT_STR_b             = "b"
	EXT_STR_dem           = "dem"
	EXT_STR_gam           = "gam"
	EXT_STR_nes           = "nes"
	EXT_STR_rom           = "rom"
	EXT_STR_sav           = "sav"
	EXT_STR_gpx           = "gpx"
	EXT_STR_kml           = "kml"
	EXT_STR_kmz           = "kmz"
	EXT_STR_log           = "log"
	EXT_STR_log1          = "log1"
	EXT_STR_crdownload    = "crdownload"
	EXT_STR_ics           = "ics"
	EXT_STR_msi           = "msi"
	EXT_STR_part          = "part"
	EXT_STR_torrent       = "torrent"
	EXT_STR_msg           = "msg"
	EXT_STR_ost           = "ost"
	EXT_STR_prf           = "prf"
	EXT_STR_pst           = "pst"
	EXT_STR_indd          = "indd"
	EXT_STR_pct           = "pct"
	EXT_STR_pdf           = "pdf"
	EXT_STR_crx           = "crx"
	EXT_STR_plugin        = "plugin"
	EXT_STR_pps           = "pps"
	EXT_STR_ppt           = "ppt"
	EXT_STR_pptm          = "pptm"
	EXT_STR_pptx          = "pptx"
	EXT_STR_bmp           = "bmp"
	EXT_STR_dds           = "dds"
	EXT_STR_gif           = "gif"
	EXT_STR_heic          = "heic"
	EXT_STR_jpeg          = "jpeg"
	EXT_STR_jpg           = "jpg"
	EXT_STR_png           = "png"
	EXT_STR_psd           = "psd"
	EXT_STR_pspimage      = "pspimage"
	EXT_STR_tga           = "tga"
	EXT_STR_thm           = "thm"
	EXT_STR_tif           = "tif"
	EXT_STR_tiff          = "tiff"
	EXT_STR_yuv           = "yuv"
	EXT_STR_cfg           = "cfg"
	EXT_STR_ini           = "ini"
	EXT_STR_xlr           = "xlr"
	EXT_STR_xls           = "xls"
	EXT_STR_xlsm          = "xlsm"
	EXT_STR_xlsx          = "xlsx"
	EXT_STR_cab           = "cab"
	EXT_STR_cpl           = "cpl"
	EXT_STR_cur           = "cur"
	EXT_STR_deskthemepack = "deskthemepack"
	EXT_STR_dll           = "dll"
	EXT_STR_dmp           = "dmp"
	EXT_STR_drv           = "drv"
	EXT_STR_icns          = "icns"
	EXT_STR_ico           = "ico"
	EXT_STR_lnk           = "lnk"
	EXT_STR_sys           = "sys"
	EXT_STR_ai            = "ai"
	EXT_STR_eps           = "eps"
	EXT_STR_ps            = "ps"
	EXT_STR_svg           = "svg"
	EXT_STR_3g2           = "3g2"
	EXT_STR_3gp           = "3gp"
	EXT_STR_asf           = "asf"
	EXT_STR_avi           = "avi"
	EXT_STR_flv           = "flv"
	EXT_STR_m4v           = "m4v"
	EXT_STR_mov           = "mov"
	EXT_STR_mp4           = "mp4"
	EXT_STR_mpg           = "mpg"
	EXT_STR_rm            = "rm"
	EXT_STR_srt           = "srt"
	EXT_STR_swf           = "swf"
	EXT_STR_vob           = "vob"
	EXT_STR_wmv           = "wmv"
	EXT_STR_asp           = "asp"
	EXT_STR_aspx          = "aspx"
	EXT_STR_cer           = "cer"
	EXT_STR_cfm           = "cfm"
	EXT_STR_csr           = "csr"
	EXT_STR_css           = "css"
	EXT_STR_dcr           = "dcr"
	EXT_STR_htm           = "htm"
	EXT_STR_html          = "html"
	EXT_STR_js            = "js"
	EXT_STR_jsp           = "jsp"
	EXT_STR_php           = "php"
	EXT_STR_rss           = "rss"
	EXT_STR_xhtml         = "xhtml"
)

const (
	FILE_Archive_File      = 1
	FILE_Audio_File        = 2
	FILE_Backup_File       = 3
	FILE_Disk_Image_File   = 4
	FILE_Log_File          = 5
	FILE_Image_File        = 6
	FILE_System_File       = 7
	FILE_Video_File        = 8
	FILE_3D_File           = 9
	FILE_CAD_File          = 10
	FILE_Data_File         = 11
	FILE_DB_File           = 12
	FILE_Developer_File    = 13
	FILE_Document_File     = 14
	FILE_Encoded_File      = 15
	FILE_Executables       = 16
	FILE_Font_File         = 17
	FILE_Game_File         = 18
	FILE_GIS_File          = 19
	FILE_Misc_File         = 20
	FILE_Outlook_File      = 21
	FILE_Page_Layout_File  = 22
	FILE_Plugin_File       = 23
	FILE_Presentation_File = 24
	FILE_Raster_Image_File = 25
	FILE_Spreadsheet_File  = 26
	FILE_Vector_Image_File = 27
	FILE_Web_File          = 28
	FILE_Settings_File     = 29
)

const (
	EXT_3dm           = 1
	EXT_3ds           = 2
	EXT_max           = 3
	EXT_obj           = 4
	EXT_7z            = 5
	EXT_cbr           = 6
	EXT_deb           = 7
	EXT_gz            = 8
	EXT_pkg           = 9
	EXT_rar           = 10
	EXT_rpm           = 11
	EXT_sitx          = 12
	EXT_tar_gz        = 13
	EXT_zip           = 14
	EXT_zipx          = 15
	EXT_aif           = 16
	EXT_iff           = 17
	EXT_m3u           = 18
	EXT_m4a           = 19
	EXT_mid           = 20
	EXT_mp3           = 21
	EXT_mpa           = 22
	EXT_wav           = 23
	EXT_wma           = 24
	EXT_bak           = 25
	EXT_old           = 26
	EXT_tmp           = 27
	EXT_dwg           = 28
	EXT_dxf           = 29
	EXT_csv           = 30
	EXT_dat           = 31
	EXT_ged           = 32
	EXT_key           = 33
	EXT_keychain      = 34
	EXT_sdf           = 35
	EXT_tar           = 36
	EXT_tax2016       = 37
	EXT_tax2018       = 38
	EXT_vcf           = 39
	EXT_xml           = 40
	EXT_accdb         = 41
	EXT_db            = 42
	EXT_dbf           = 43
	EXT_mdb           = 44
	EXT_pdb           = 45
	EXT_sql           = 46
	EXT_sqlite_wal    = 47
	EXT_c             = 48
	EXT_class         = 49
	EXT_cpp           = 50
	EXT_cs            = 51
	EXT_dtd           = 52
	EXT_fla           = 53
	EXT_h             = 54
	EXT_java          = 55
	EXT_lua           = 56
	EXT_m             = 57
	EXT_mpp           = 58
	EXT_pl            = 59
	EXT_py            = 60
	EXT_sh            = 61
	EXT_sln           = 62
	EXT_swift         = 63
	EXT_vb            = 64
	EXT_vcxproj       = 65
	EXT_xcodeproj     = 66
	EXT_bin           = 67
	EXT_cue           = 68
	EXT_dmg           = 69
	EXT_iso           = 70
	EXT_mdf           = 71
	EXT_toast         = 72
	EXT_vcd           = 73
	EXT_doc           = 74
	EXT_docx          = 75
	EXT_odt           = 76
	EXT_pages         = 77
	EXT_rtf           = 78
	EXT_tex           = 79
	EXT_txt           = 80
	EXT_vsd           = 81
	EXT_wpd           = 82
	EXT_wps           = 83
	EXT_hqx           = 84
	EXT_mim           = 85
	EXT_uue           = 86
	EXT_apk           = 87
	EXT_app           = 88
	EXT_bat           = 89
	EXT_cgi           = 90
	EXT_com           = 91
	EXT_exe           = 92
	EXT_gadget        = 93
	EXT_jar           = 94
	EXT_wsf           = 95
	EXT_fnt           = 96
	EXT_fon           = 97
	EXT_otf           = 98
	EXT_ttf           = 99
	EXT_b             = 100
	EXT_dem           = 101
	EXT_gam           = 102
	EXT_nes           = 103
	EXT_rom           = 104
	EXT_sav           = 105
	EXT_gpx           = 106
	EXT_kml           = 107
	EXT_kmz           = 108
	EXT_log           = 109
	EXT_log1          = 110
	EXT_crdownload    = 111
	EXT_ics           = 112
	EXT_msi           = 113
	EXT_part          = 114
	EXT_torrent       = 115
	EXT_msg           = 116
	EXT_ost           = 117
	EXT_prf           = 118
	EXT_pst           = 119
	EXT_indd          = 120
	EXT_pct           = 121
	EXT_pdf           = 122
	EXT_crx           = 123
	EXT_plugin        = 124
	EXT_pps           = 125
	EXT_ppt           = 126
	EXT_pptm          = 127
	EXT_pptx          = 128
	EXT_bmp           = 129
	EXT_dds           = 130
	EXT_gif           = 131
	EXT_heic          = 132
	EXT_jpeg          = 133
	EXT_jpg           = 134
	EXT_png           = 135
	EXT_psd           = 136
	EXT_pspimage      = 137
	EXT_tga           = 138
	EXT_thm           = 139
	EXT_tif           = 140
	EXT_tiff          = 141
	EXT_yuv           = 142
	EXT_cfg           = 143
	EXT_ini           = 144
	EXT_xlr           = 145
	EXT_xls           = 146
	EXT_xlsm          = 147
	EXT_xlsx          = 148
	EXT_cab           = 149
	EXT_cpl           = 150
	EXT_cur           = 151
	EXT_deskthemepack = 152
	EXT_dll           = 153
	EXT_dmp           = 154
	EXT_drv           = 155
	EXT_icns          = 156
	EXT_ico           = 157
	EXT_lnk           = 158
	EXT_sys           = 159
	EXT_ai            = 160
	EXT_eps           = 161
	EXT_ps            = 162
	EXT_svg           = 163
	EXT_3g2           = 164
	EXT_3gp           = 165
	EXT_asf           = 166
	EXT_avi           = 167
	EXT_flv           = 168
	EXT_m4v           = 169
	EXT_mov           = 170
	EXT_mp4           = 171
	EXT_mpg           = 172
	EXT_rm            = 173
	EXT_srt           = 174
	EXT_swf           = 175
	EXT_vob           = 176
	EXT_wmv           = 177
	EXT_asp           = 178
	EXT_aspx          = 179
	EXT_cer           = 180
	EXT_cfm           = 181
	EXT_csr           = 182
	EXT_css           = 183
	EXT_dcr           = 184
	EXT_htm           = 185
	EXT_html          = 186
	EXT_js            = 187
	EXT_jsp           = 188
	EXT_php           = 189
	EXT_rss           = 190
	EXT_xhtml         = 191
)

var Ext = map[int]string{
	EXT_3dm:           EXT_STR_3dm,
	EXT_3ds:           EXT_STR_3ds,
	EXT_max:           EXT_STR_max,
	EXT_obj:           EXT_STR_obj,
	EXT_7z:            EXT_STR_7z,
	EXT_cbr:           EXT_STR_cbr,
	EXT_deb:           EXT_STR_deb,
	EXT_gz:            EXT_STR_gz,
	EXT_pkg:           EXT_STR_pkg,
	EXT_rar:           EXT_STR_rar,
	EXT_rpm:           EXT_STR_rpm,
	EXT_sitx:          EXT_STR_sitx,
	EXT_tar_gz:        EXT_STR_tar_gz,
	EXT_zip:           EXT_STR_zip,
	EXT_zipx:          EXT_STR_zipx,
	EXT_aif:           EXT_STR_aif,
	EXT_iff:           EXT_STR_iff,
	EXT_m3u:           EXT_STR_m3u,
	EXT_m4a:           EXT_STR_m4a,
	EXT_mid:           EXT_STR_mid,
	EXT_mp3:           EXT_STR_mp3,
	EXT_mpa:           EXT_STR_mpa,
	EXT_wav:           EXT_STR_wav,
	EXT_wma:           EXT_STR_wma,
	EXT_bak:           EXT_STR_bak,
	EXT_old:           EXT_STR_old,
	EXT_tmp:           EXT_STR_tmp,
	EXT_dwg:           EXT_STR_dwg,
	EXT_dxf:           EXT_STR_dxf,
	EXT_csv:           EXT_STR_csv,
	EXT_dat:           EXT_STR_dat,
	EXT_ged:           EXT_STR_ged,
	EXT_key:           EXT_STR_key,
	EXT_keychain:      EXT_STR_keychain,
	EXT_sdf:           EXT_STR_sdf,
	EXT_tar:           EXT_STR_tar,
	EXT_tax2016:       EXT_STR_tax2016,
	EXT_tax2018:       EXT_STR_tax2018,
	EXT_vcf:           EXT_STR_vcf,
	EXT_xml:           EXT_STR_xml,
	EXT_accdb:         EXT_STR_accdb,
	EXT_db:            EXT_STR_db,
	EXT_dbf:           EXT_STR_dbf,
	EXT_mdb:           EXT_STR_mdb,
	EXT_pdb:           EXT_STR_pdb,
	EXT_sql:           EXT_STR_sql,
	EXT_sqlite_wal:    EXT_STR_sqlite_wal,
	EXT_c:             EXT_STR_c,
	EXT_class:         EXT_STR_class,
	EXT_cpp:           EXT_STR_cpp,
	EXT_cs:            EXT_STR_cs,
	EXT_dtd:           EXT_STR_dtd,
	EXT_fla:           EXT_STR_fla,
	EXT_h:             EXT_STR_h,
	EXT_java:          EXT_STR_java,
	EXT_lua:           EXT_STR_lua,
	EXT_m:             EXT_STR_m,
	EXT_mpp:           EXT_STR_mpp,
	EXT_pl:            EXT_STR_pl,
	EXT_py:            EXT_STR_py,
	EXT_sh:            EXT_STR_sh,
	EXT_sln:           EXT_STR_sln,
	EXT_swift:         EXT_STR_swift,
	EXT_vb:            EXT_STR_vb,
	EXT_vcxproj:       EXT_STR_vcxproj,
	EXT_xcodeproj:     EXT_STR_xcodeproj,
	EXT_bin:           EXT_STR_bin,
	EXT_cue:           EXT_STR_cue,
	EXT_dmg:           EXT_STR_dmg,
	EXT_iso:           EXT_STR_iso,
	EXT_mdf:           EXT_STR_mdf,
	EXT_toast:         EXT_STR_toast,
	EXT_vcd:           EXT_STR_vcd,
	EXT_doc:           EXT_STR_doc,
	EXT_docx:          EXT_STR_docx,
	EXT_odt:           EXT_STR_odt,
	EXT_pages:         EXT_STR_pages,
	EXT_rtf:           EXT_STR_rtf,
	EXT_tex:           EXT_STR_tex,
	EXT_txt:           EXT_STR_txt,
	EXT_vsd:           EXT_STR_vsd,
	EXT_wpd:           EXT_STR_wpd,
	EXT_wps:           EXT_STR_wps,
	EXT_hqx:           EXT_STR_hqx,
	EXT_mim:           EXT_STR_mim,
	EXT_uue:           EXT_STR_uue,
	EXT_apk:           EXT_STR_apk,
	EXT_app:           EXT_STR_app,
	EXT_bat:           EXT_STR_bat,
	EXT_cgi:           EXT_STR_cgi,
	EXT_com:           EXT_STR_com,
	EXT_exe:           EXT_STR_exe,
	EXT_gadget:        EXT_STR_gadget,
	EXT_jar:           EXT_STR_jar,
	EXT_wsf:           EXT_STR_wsf,
	EXT_fnt:           EXT_STR_fnt,
	EXT_fon:           EXT_STR_fon,
	EXT_otf:           EXT_STR_otf,
	EXT_ttf:           EXT_STR_ttf,
	EXT_b:             EXT_STR_b,
	EXT_dem:           EXT_STR_dem,
	EXT_gam:           EXT_STR_gam,
	EXT_nes:           EXT_STR_nes,
	EXT_rom:           EXT_STR_rom,
	EXT_sav:           EXT_STR_sav,
	EXT_gpx:           EXT_STR_gpx,
	EXT_kml:           EXT_STR_kml,
	EXT_kmz:           EXT_STR_kmz,
	EXT_log:           EXT_STR_log,
	EXT_log1:          EXT_STR_log1,
	EXT_crdownload:    EXT_STR_crdownload,
	EXT_ics:           EXT_STR_ics,
	EXT_msi:           EXT_STR_msi,
	EXT_part:          EXT_STR_part,
	EXT_torrent:       EXT_STR_torrent,
	EXT_msg:           EXT_STR_msg,
	EXT_ost:           EXT_STR_ost,
	EXT_prf:           EXT_STR_prf,
	EXT_pst:           EXT_STR_pst,
	EXT_indd:          EXT_STR_indd,
	EXT_pct:           EXT_STR_pct,
	EXT_pdf:           EXT_STR_pdf,
	EXT_crx:           EXT_STR_crx,
	EXT_plugin:        EXT_STR_plugin,
	EXT_pps:           EXT_STR_pps,
	EXT_ppt:           EXT_STR_ppt,
	EXT_pptm:          EXT_STR_pptm,
	EXT_pptx:          EXT_STR_pptx,
	EXT_bmp:           EXT_STR_bmp,
	EXT_dds:           EXT_STR_dds,
	EXT_gif:           EXT_STR_gif,
	EXT_heic:          EXT_STR_heic,
	EXT_jpeg:          EXT_STR_jpeg,
	EXT_jpg:           EXT_STR_jpg,
	EXT_png:           EXT_STR_png,
	EXT_psd:           EXT_STR_psd,
	EXT_pspimage:      EXT_STR_pspimage,
	EXT_tga:           EXT_STR_tga,
	EXT_thm:           EXT_STR_thm,
	EXT_tif:           EXT_STR_tif,
	EXT_tiff:          EXT_STR_tiff,
	EXT_yuv:           EXT_STR_yuv,
	EXT_cfg:           EXT_STR_cfg,
	EXT_ini:           EXT_STR_ini,
	EXT_xlr:           EXT_STR_xlr,
	EXT_xls:           EXT_STR_xls,
	EXT_xlsm:          EXT_STR_xlsm,
	EXT_xlsx:          EXT_STR_xlsx,
	EXT_cab:           EXT_STR_cab,
	EXT_cpl:           EXT_STR_cpl,
	EXT_cur:           EXT_STR_cur,
	EXT_deskthemepack: EXT_STR_deskthemepack,
	EXT_dll:           EXT_STR_dll,
	EXT_dmp:           EXT_STR_dmp,
	EXT_drv:           EXT_STR_drv,
	EXT_icns:          EXT_STR_icns,
	EXT_ico:           EXT_STR_ico,
	EXT_lnk:           EXT_STR_lnk,
	EXT_sys:           EXT_STR_sys,
	EXT_ai:            EXT_STR_ai,
	EXT_eps:           EXT_STR_eps,
	EXT_ps:            EXT_STR_ps,
	EXT_svg:           EXT_STR_svg,
	EXT_3g2:           EXT_STR_3g2,
	EXT_3gp:           EXT_STR_3gp,
	EXT_asf:           EXT_STR_asf,
	EXT_avi:           EXT_STR_avi,
	EXT_flv:           EXT_STR_flv,
	EXT_m4v:           EXT_STR_m4v,
	EXT_mov:           EXT_STR_mov,
	EXT_mp4:           EXT_STR_mp4,
	EXT_mpg:           EXT_STR_mpg,
	EXT_rm:            EXT_STR_rm,
	EXT_srt:           EXT_STR_srt,
	EXT_swf:           EXT_STR_swf,
	EXT_vob:           EXT_STR_vob,
	EXT_wmv:           EXT_STR_wmv,
	EXT_asp:           EXT_STR_asp,
	EXT_aspx:          EXT_STR_aspx,
	EXT_cer:           EXT_STR_cer,
	EXT_cfm:           EXT_STR_cfm,
	EXT_csr:           EXT_STR_csr,
	EXT_css:           EXT_STR_css,
	EXT_dcr:           EXT_STR_dcr,
	EXT_htm:           EXT_STR_htm,
	EXT_html:          EXT_STR_html,
	EXT_js:            EXT_STR_js,
	EXT_jsp:           EXT_STR_jsp,
	EXT_php:           EXT_STR_php,
	EXT_rss:           EXT_STR_rss,
	EXT_xhtml:         EXT_STR_xhtml,
}

var ExtId = map[string]int{
	EXT_STR_3dm:           EXT_3dm,
	EXT_STR_3ds:           EXT_3ds,
	EXT_STR_max:           EXT_max,
	EXT_STR_obj:           EXT_obj,
	EXT_STR_7z:            EXT_7z,
	EXT_STR_cbr:           EXT_cbr,
	EXT_STR_deb:           EXT_deb,
	EXT_STR_gz:            EXT_gz,
	EXT_STR_pkg:           EXT_pkg,
	EXT_STR_rar:           EXT_rar,
	EXT_STR_rpm:           EXT_rpm,
	EXT_STR_sitx:          EXT_sitx,
	EXT_STR_tar_gz:        EXT_tar_gz,
	EXT_STR_zip:           EXT_zip,
	EXT_STR_zipx:          EXT_zipx,
	EXT_STR_aif:           EXT_aif,
	EXT_STR_iff:           EXT_iff,
	EXT_STR_m3u:           EXT_m3u,
	EXT_STR_m4a:           EXT_m4a,
	EXT_STR_mid:           EXT_mid,
	EXT_STR_mp3:           EXT_mp3,
	EXT_STR_mpa:           EXT_mpa,
	EXT_STR_wav:           EXT_wav,
	EXT_STR_wma:           EXT_wma,
	EXT_STR_bak:           EXT_bak,
	EXT_STR_old:           EXT_old,
	EXT_STR_tmp:           EXT_tmp,
	EXT_STR_dwg:           EXT_dwg,
	EXT_STR_dxf:           EXT_dxf,
	EXT_STR_csv:           EXT_csv,
	EXT_STR_dat:           EXT_dat,
	EXT_STR_ged:           EXT_ged,
	EXT_STR_key:           EXT_key,
	EXT_STR_keychain:      EXT_keychain,
	EXT_STR_sdf:           EXT_sdf,
	EXT_STR_tar:           EXT_tar,
	EXT_STR_tax2016:       EXT_tax2016,
	EXT_STR_tax2018:       EXT_tax2018,
	EXT_STR_vcf:           EXT_vcf,
	EXT_STR_xml:           EXT_xml,
	EXT_STR_accdb:         EXT_accdb,
	EXT_STR_db:            EXT_db,
	EXT_STR_dbf:           EXT_dbf,
	EXT_STR_mdb:           EXT_mdb,
	EXT_STR_pdb:           EXT_pdb,
	EXT_STR_sql:           EXT_sql,
	EXT_STR_sqlite_wal:    EXT_sqlite_wal,
	EXT_STR_c:             EXT_c,
	EXT_STR_class:         EXT_class,
	EXT_STR_cpp:           EXT_cpp,
	EXT_STR_cs:            EXT_cs,
	EXT_STR_dtd:           EXT_dtd,
	EXT_STR_fla:           EXT_fla,
	EXT_STR_h:             EXT_h,
	EXT_STR_java:          EXT_java,
	EXT_STR_lua:           EXT_lua,
	EXT_STR_m:             EXT_m,
	EXT_STR_mpp:           EXT_mpp,
	EXT_STR_pl:            EXT_pl,
	EXT_STR_py:            EXT_py,
	EXT_STR_sh:            EXT_sh,
	EXT_STR_sln:           EXT_sln,
	EXT_STR_swift:         EXT_swift,
	EXT_STR_vb:            EXT_vb,
	EXT_STR_vcxproj:       EXT_vcxproj,
	EXT_STR_xcodeproj:     EXT_xcodeproj,
	EXT_STR_bin:           EXT_bin,
	EXT_STR_cue:           EXT_cue,
	EXT_STR_dmg:           EXT_dmg,
	EXT_STR_iso:           EXT_iso,
	EXT_STR_mdf:           EXT_mdf,
	EXT_STR_toast:         EXT_toast,
	EXT_STR_vcd:           EXT_vcd,
	EXT_STR_doc:           EXT_doc,
	EXT_STR_docx:          EXT_docx,
	EXT_STR_odt:           EXT_odt,
	EXT_STR_pages:         EXT_pages,
	EXT_STR_rtf:           EXT_rtf,
	EXT_STR_tex:           EXT_tex,
	EXT_STR_txt:           EXT_txt,
	EXT_STR_vsd:           EXT_vsd,
	EXT_STR_wpd:           EXT_wpd,
	EXT_STR_wps:           EXT_wps,
	EXT_STR_hqx:           EXT_hqx,
	EXT_STR_mim:           EXT_mim,
	EXT_STR_uue:           EXT_uue,
	EXT_STR_apk:           EXT_apk,
	EXT_STR_app:           EXT_app,
	EXT_STR_bat:           EXT_bat,
	EXT_STR_cgi:           EXT_cgi,
	EXT_STR_com:           EXT_com,
	EXT_STR_exe:           EXT_exe,
	EXT_STR_gadget:        EXT_gadget,
	EXT_STR_jar:           EXT_jar,
	EXT_STR_wsf:           EXT_wsf,
	EXT_STR_fnt:           EXT_fnt,
	EXT_STR_fon:           EXT_fon,
	EXT_STR_otf:           EXT_otf,
	EXT_STR_ttf:           EXT_ttf,
	EXT_STR_b:             EXT_b,
	EXT_STR_dem:           EXT_dem,
	EXT_STR_gam:           EXT_gam,
	EXT_STR_nes:           EXT_nes,
	EXT_STR_rom:           EXT_rom,
	EXT_STR_sav:           EXT_sav,
	EXT_STR_gpx:           EXT_gpx,
	EXT_STR_kml:           EXT_kml,
	EXT_STR_kmz:           EXT_kmz,
	EXT_STR_log:           EXT_log,
	EXT_STR_log1:          EXT_log1,
	EXT_STR_crdownload:    EXT_crdownload,
	EXT_STR_ics:           EXT_ics,
	EXT_STR_msi:           EXT_msi,
	EXT_STR_part:          EXT_part,
	EXT_STR_torrent:       EXT_torrent,
	EXT_STR_msg:           EXT_msg,
	EXT_STR_ost:           EXT_ost,
	EXT_STR_prf:           EXT_prf,
	EXT_STR_pst:           EXT_pst,
	EXT_STR_indd:          EXT_indd,
	EXT_STR_pct:           EXT_pct,
	EXT_STR_pdf:           EXT_pdf,
	EXT_STR_crx:           EXT_crx,
	EXT_STR_plugin:        EXT_plugin,
	EXT_STR_pps:           EXT_pps,
	EXT_STR_ppt:           EXT_ppt,
	EXT_STR_pptm:          EXT_pptm,
	EXT_STR_pptx:          EXT_pptx,
	EXT_STR_bmp:           EXT_bmp,
	EXT_STR_dds:           EXT_dds,
	EXT_STR_gif:           EXT_gif,
	EXT_STR_heic:          EXT_heic,
	EXT_STR_jpeg:          EXT_jpeg,
	EXT_STR_jpg:           EXT_jpg,
	EXT_STR_png:           EXT_png,
	EXT_STR_psd:           EXT_psd,
	EXT_STR_pspimage:      EXT_pspimage,
	EXT_STR_tga:           EXT_tga,
	EXT_STR_thm:           EXT_thm,
	EXT_STR_tif:           EXT_tif,
	EXT_STR_tiff:          EXT_tiff,
	EXT_STR_yuv:           EXT_yuv,
	EXT_STR_cfg:           EXT_cfg,
	EXT_STR_ini:           EXT_ini,
	EXT_STR_xlr:           EXT_xlr,
	EXT_STR_xls:           EXT_xls,
	EXT_STR_xlsm:          EXT_xlsm,
	EXT_STR_xlsx:          EXT_xlsx,
	EXT_STR_cab:           EXT_cab,
	EXT_STR_cpl:           EXT_cpl,
	EXT_STR_cur:           EXT_cur,
	EXT_STR_deskthemepack: EXT_deskthemepack,
	EXT_STR_dll:           EXT_dll,
	EXT_STR_dmp:           EXT_dmp,
	EXT_STR_drv:           EXT_drv,
	EXT_STR_icns:          EXT_icns,
	EXT_STR_ico:           EXT_ico,
	EXT_STR_lnk:           EXT_lnk,
	EXT_STR_sys:           EXT_sys,
	EXT_STR_ai:            EXT_ai,
	EXT_STR_eps:           EXT_eps,
	EXT_STR_ps:            EXT_ps,
	EXT_STR_svg:           EXT_svg,
	EXT_STR_3g2:           EXT_3g2,
	EXT_STR_3gp:           EXT_3gp,
	EXT_STR_asf:           EXT_asf,
	EXT_STR_avi:           EXT_avi,
	EXT_STR_flv:           EXT_flv,
	EXT_STR_m4v:           EXT_m4v,
	EXT_STR_mov:           EXT_mov,
	EXT_STR_mp4:           EXT_mp4,
	EXT_STR_mpg:           EXT_mpg,
	EXT_STR_rm:            EXT_rm,
	EXT_STR_srt:           EXT_srt,
	EXT_STR_swf:           EXT_swf,
	EXT_STR_vob:           EXT_vob,
	EXT_STR_wmv:           EXT_wmv,
	EXT_STR_asp:           EXT_asp,
	EXT_STR_aspx:          EXT_aspx,
	EXT_STR_cer:           EXT_cer,
	EXT_STR_cfm:           EXT_cfm,
	EXT_STR_csr:           EXT_csr,
	EXT_STR_css:           EXT_css,
	EXT_STR_dcr:           EXT_dcr,
	EXT_STR_htm:           EXT_htm,
	EXT_STR_html:          EXT_html,
	EXT_STR_js:            EXT_js,
	EXT_STR_jsp:           EXT_jsp,
	EXT_STR_php:           EXT_php,
	EXT_STR_rss:           EXT_rss,
	EXT_STR_xhtml:         EXT_xhtml,
}

const (
	FILE_STR_Archive_File      = "Archive Files"
	FILE_STR_Audio_File        = "Audio Files"
	FILE_STR_Backup_File       = "Backup Files"
	FILE_STR_Disk_Image_File   = "Disk Image Files"
	FILE_STR_Log_File          = "Log Files"
	FILE_STR_Image_File        = "Image Files"
	FILE_STR_System_File       = "System Files"
	FILE_STR_Video_File        = "Video Files"
	FILE_STR_3D_File           = "3D Files"
	FILE_STR_CAD_File          = "CAD Files"
	FILE_STR_Data_File         = "Data Files"
	FILE_STR_DB_File           = "DB Files"
	FILE_STR_Developer_File    = "Developer Files"
	FILE_STR_Document_File     = "Document Files"
	FILE_STR_Encoded_File      = "Encoded Files"
	FILE_STR_Executables       = "Executables"
	FILE_STR_Font_File         = "Font Files"
	FILE_STR_Game_File         = "Game Files"
	FILE_STR_GIS_File          = "GIS Files"
	FILE_STR_Misc_File         = "Misc Files"
	FILE_STR_Outlook_File      = "Outlook Files"
	FILE_STR_Page_Layout_File  = "Page Layout Files"
	FILE_STR_Plugin_File       = "Plugin Files"
	FILE_STR_Presentation_File = "Presentation Files"
	FILE_STR_Raster_Image_File = "Raster Image Files"
	FILE_STR_Spreadsheet_File  = "Spreadsheet Files"
	FILE_STR_Vector_Image_File = "Vector Image Files"
	FILE_STR_Web_File          = "Web Files"
	FILE_STR_Settings_File     = "Settings Files"
)

var FileType = map[int]string{
	FILE_Archive_File:      FILE_STR_Archive_File,
	FILE_Audio_File:        FILE_STR_Audio_File,
	FILE_Backup_File:       FILE_STR_Backup_File,
	FILE_Disk_Image_File:   FILE_STR_Disk_Image_File,
	FILE_Log_File:          FILE_STR_Log_File,
	FILE_Image_File:        FILE_STR_Image_File,
	FILE_System_File:       FILE_STR_System_File,
	FILE_Video_File:        FILE_STR_Video_File,
	FILE_3D_File:           FILE_STR_3D_File,
	FILE_CAD_File:          FILE_STR_CAD_File,
	FILE_Data_File:         FILE_STR_Data_File,
	FILE_DB_File:           FILE_STR_DB_File,
	FILE_Developer_File:    FILE_STR_Developer_File,
	FILE_Document_File:     FILE_STR_Document_File,
	FILE_Encoded_File:      FILE_STR_Encoded_File,
	FILE_Executables:       FILE_STR_Executables,
	FILE_Font_File:         FILE_STR_Font_File,
	FILE_Game_File:         FILE_STR_Game_File,
	FILE_GIS_File:          FILE_STR_GIS_File,
	FILE_Misc_File:         FILE_STR_Misc_File,
	FILE_Outlook_File:      FILE_STR_Outlook_File,
	FILE_Page_Layout_File:  FILE_STR_Page_Layout_File,
	FILE_Plugin_File:       FILE_STR_Plugin_File,
	FILE_Presentation_File: FILE_STR_Presentation_File,
	FILE_Raster_Image_File: FILE_STR_Raster_Image_File,
	FILE_Spreadsheet_File:  FILE_STR_Spreadsheet_File,
	FILE_Vector_Image_File: FILE_STR_Vector_Image_File,
	FILE_Web_File:          FILE_STR_Web_File,
	FILE_Settings_File:     FILE_STR_Settings_File,
}

//list of supported file type
var Filelist = []int{
	FILE_Archive_File,
	FILE_Audio_File,
	FILE_Backup_File,
	FILE_Disk_Image_File,
	FILE_Log_File,
	FILE_Image_File,
	FILE_System_File,
	FILE_Video_File,
	FILE_3D_File,
	FILE_CAD_File,
	FILE_Data_File,
	FILE_DB_File,
	FILE_Developer_File,
	FILE_Document_File,
	FILE_Encoded_File,
	FILE_Executables,
	FILE_Font_File,
	FILE_Game_File,
	FILE_GIS_File,
	FILE_Misc_File,
	FILE_Outlook_File,
	FILE_Page_Layout_File,
	FILE_Plugin_File,
	FILE_Presentation_File,
	FILE_Raster_Image_File,
	FILE_Spreadsheet_File,
	FILE_Vector_Image_File,
	FILE_Web_File,
	FILE_Settings_File,
}

var FileExts = map[int][]string{
	FILE_3D_File: []string{EXT_STR_3dm,
		EXT_STR_3ds,
		EXT_STR_max,
		EXT_STR_obj,
	},
	FILE_CAD_File: []string{EXT_STR_dwg,
		EXT_STR_dxf,
	},
	FILE_Data_File: []string{EXT_STR_csv,
		EXT_STR_dat,
		EXT_STR_ged,
		EXT_STR_key,
		EXT_STR_keychain,
		EXT_STR_sdf,
		EXT_STR_tar,
		EXT_STR_tax2016,
		EXT_STR_tax2018,
		EXT_STR_vcf,
		EXT_STR_xml,
	},
	FILE_DB_File: []string{EXT_STR_accdb,
		EXT_STR_db,
		EXT_STR_dbf,
		EXT_STR_mdb,
		EXT_STR_pdb,
		EXT_STR_sql,
		EXT_STR_sqlite_wal,
	},
	FILE_Developer_File: []string{EXT_STR_c,
		EXT_STR_class,
		EXT_STR_cpp,
		EXT_STR_cs,
		EXT_STR_dtd,
		EXT_STR_fla,
		EXT_STR_h,
		EXT_STR_java,
		EXT_STR_lua,
		EXT_STR_m,
		EXT_STR_mpp,
		EXT_STR_pl,
		EXT_STR_py,
		EXT_STR_sh,
		EXT_STR_sln,
		EXT_STR_swift,
		EXT_STR_vb,
		EXT_STR_vcxproj,
		EXT_STR_xcodeproj,
	},
	FILE_Document_File: []string{EXT_STR_doc,
		EXT_STR_docx,
		EXT_STR_odt,
		EXT_STR_pages,
		EXT_STR_rtf,
		EXT_STR_tex,
		EXT_STR_txt,
		EXT_STR_vsd,
		EXT_STR_wpd,
		EXT_STR_wps,
	},
	FILE_Encoded_File: []string{EXT_STR_hqx,
		EXT_STR_mim,
		EXT_STR_uue,
	},
	FILE_Executables: []string{EXT_STR_apk,
		EXT_STR_app,
		EXT_STR_bat,
		EXT_STR_cgi,
		EXT_STR_com,
		EXT_STR_exe,
		EXT_STR_gadget,
		EXT_STR_jar,
		EXT_STR_wsf,
	},
	FILE_Font_File: []string{EXT_STR_fnt,
		EXT_STR_fon,
		EXT_STR_otf,
		EXT_STR_ttf,
	},
	FILE_Game_File: []string{EXT_STR_b,
		EXT_STR_dem,
		EXT_STR_gam,
		EXT_STR_nes,
		EXT_STR_rom,
		EXT_STR_sav,
	},
	FILE_GIS_File: []string{EXT_STR_gpx,
		EXT_STR_kml,
		EXT_STR_kmz,
	},
	FILE_Misc_File: []string{EXT_STR_crdownload,
		EXT_STR_ics,
		EXT_STR_msi,
		EXT_STR_part,
		EXT_STR_torrent,
	},
	FILE_Outlook_File: []string{EXT_STR_msg,
		EXT_STR_ost,
		EXT_STR_prf,
		EXT_STR_pst,
	},
	FILE_Page_Layout_File: []string{EXT_STR_indd,
		EXT_STR_pct,
		EXT_STR_pdf,
	},
	FILE_Plugin_File: []string{EXT_STR_crx,
		EXT_STR_plugin,
	},
	FILE_Presentation_File: []string{EXT_STR_pps,
		EXT_STR_ppt,
		EXT_STR_pptm,
		EXT_STR_pptx,
	},
	FILE_Raster_Image_File: []string{EXT_STR_dds,
		EXT_STR_heic,
		EXT_STR_psd,
		EXT_STR_tga,
		EXT_STR_yuv,
	},
	FILE_Settings_File: []string{EXT_STR_cfg,
		EXT_STR_ini,
	},
	FILE_Spreadsheet_File: []string{EXT_STR_xlr,
		EXT_STR_xls,
		EXT_STR_xlsm,
		EXT_STR_xlsx,
	},
	FILE_Vector_Image_File: []string{EXT_STR_ai,
		EXT_STR_eps,
		EXT_STR_ps,
		EXT_STR_svg,
	},
	FILE_Web_File: []string{EXT_STR_asp,
		EXT_STR_aspx,
		EXT_STR_cer,
		EXT_STR_cfm,
		EXT_STR_csr,
		EXT_STR_css,
		EXT_STR_dcr,
		EXT_STR_htm,
		EXT_STR_html,
		EXT_STR_js,
		EXT_STR_jsp,
		EXT_STR_php,
		EXT_STR_rss,
		EXT_STR_xhtml,
	},
	FILE_Archive_File: []string{
		EXT_STR_7z,
		EXT_STR_cbr,
		EXT_STR_deb,
		EXT_STR_gz,
		EXT_STR_pkg,
		EXT_STR_rar,
		EXT_STR_rpm,
		EXT_STR_sitx,
		EXT_STR_tar_gz,
		EXT_STR_zip,
		EXT_STR_zipx,
	},
	FILE_Audio_File: []string{
		EXT_STR_aif,
		EXT_STR_iff,
		EXT_STR_m3u,
		EXT_STR_m4a,
		EXT_STR_mid,
		EXT_STR_mp3,
		EXT_STR_mpa,
		EXT_STR_wav,
		EXT_STR_wma,
	},
	FILE_Backup_File: []string{
		EXT_STR_bak,
		EXT_STR_old,
		EXT_STR_tmp,
	},
	FILE_Disk_Image_File: []string{
		EXT_STR_bin,
		EXT_STR_cue,
		EXT_STR_dmg,
		EXT_STR_iso,
		EXT_STR_mdf,
		EXT_STR_toast,
		EXT_STR_vcd,
	},
	FILE_Log_File: []string{
		EXT_STR_log,
		EXT_STR_log1,
	},
	FILE_Image_File: []string{
		EXT_STR_bmp,
		EXT_STR_gif,
		EXT_STR_jpeg,
		EXT_STR_jpg,
		EXT_STR_png,
		EXT_STR_pspimage,
		EXT_STR_thm,
		EXT_STR_tif,
		EXT_STR_tiff,
	},
	FILE_System_File: []string{
		EXT_STR_cab,
		EXT_STR_cpl,
		EXT_STR_cur,
		EXT_STR_deskthemepack,
		EXT_STR_dll,
		EXT_STR_dmp,
		EXT_STR_drv,
		EXT_STR_icns,
		EXT_STR_ico,
		EXT_STR_lnk,
		EXT_STR_sys,
	},
	FILE_Video_File: []string{
		EXT_STR_3g2,
		EXT_STR_3gp,
		EXT_STR_asf,
		EXT_STR_avi,
		EXT_STR_flv,
		EXT_STR_m4v,
		EXT_STR_mov,
		EXT_STR_mp4,
		EXT_STR_mpg,
		EXT_STR_rm,
		EXT_STR_srt,
		EXT_STR_swf,
		EXT_STR_vob,
		EXT_STR_wmv,
	},
}

//month
const (
	January   = 1
	February  = 2
	March     = 3
	April     = 4
	May       = 5
	June      = 6
	July      = 7
	August    = 8
	September = 9
	October   = 10
	November  = 11
	December  = 12
)
const (
	JanuaryStr   = "Jan"
	FebruaryStr  = "Feb"
	MarchStr     = "Mar"
	AprilStr     = "Apr"
	MayStr       = "May"
	JuneStr      = "Jun"
	JulyStr      = "Jul"
	AugustStr    = "Aug"
	SeptemberStr = "Sep"
	OctoberStr   = "Oct"
	NovemberStr  = "Nov"
	DecemberStr  = "Dec"
)

var MonthMap = map[int]string{
	January:   JanuaryStr,
	February:  FebruaryStr,
	March:     MarchStr,
	April:     AprilStr,
	May:       MayStr,
	June:      JuneStr,
	July:      JulyStr,
	August:    AugustStr,
	September: SeptemberStr,
	October:   OctoberStr,
	November:  NovemberStr,
	December:  DecemberStr,
}
