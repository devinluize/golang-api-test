package repositories

import (
	"NewProjectTestingApi/entities"
	"context"
	"database/sql"
	"gorm.io/gorm"
)

func LoggingIfError(ctx context.Context, db *gorm.DB, params entities.LogbookInsertParams) int {
	tx := db.Begin()
	defer tx.Rollback()
	var logSysNo int
	query := `
    exec api.uspg_logbook_insert 
        @LOG_SYS_NO=@p1 output,
        @API_NAME=@p2,
        @API_METHOD=@p3,
        @API_URL=@p4,
        @API_PARAMETERS=@p5,
        @API_START=@p6,
        @API_FINISH=@p7,
        @API_STATUS=@p8,
        @REQUEST_HEADER=@p9,
        @REQUEST_BODY=@p10,
        @RESPONSE_HEADER=@p11,
        @RESPONSE_BODY=@p12,
        @ROWS_AFFECTED=@p13,
        @SUCCESS=@p14,
        @MESSAGE=@p15,
        @PROCESS_FROM=@p16,
        @PROCESS_BY=@p17
`
	db.Raw(query,
		sql.Out{Dest: &logSysNo},          // @p1 - Output parameter
		"/DMSMobile/ViewListProspectData", // @p2
		"POST",                            // @p3
		"/DMSMobile/ViewListProspectData", // @p4
		`{"strSalesRep":"105196","strCompany":"131302","strProspectName":"","strModel":"","strBrand":"","strVariant":"","strStage":null,"intPageNumber":1}`, // @p5
		"2024-04-03 08:45:32.377", // @p6
		"2024-04-03 08:45:32.470", // @p7
		"200",                     // @p8
		`{"Accept":["application/json"],"Accept-Encoding":["gzip"],"Authorization":["Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1lIjoiQWRtaW5ETVNNb2JpbGUiLCJqdGkiOiJiYjM1MjA5ZS1hNzY4LTQxMDQtYmM0OC0wNWI2ZTkwZTBjOTciLCJodHRwOi8vc2NoZW1hcy5taWNyb3NvZnQuY29tL3dzLzIwMDgvMDYvaWRlbnRpdHkvY2xhaW1zL3JvbGUiOiJETVNNb2JpbGUiLCJleHAiOjE3MTI1NDM1OTYsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3Q6NjE5NTUiLCJhdWQiOiJodHRwOi8vbG9jYWxob3N0OjQyMDAifQ.2wrZdGiPUza15fub3exVqiH0f3oUtPcbuw-NbGVQg7c"],"Connection":["upgrade"],"Content-Length":["131"],"Content-Type":["application/json; charset=utf-8"],"Host":["dms-api.indomobil.co.id"],"User-Agent":["Dart/2.19 (dart:io)"],"X-Real-IP":["114.10.5.46"],"X-Forwarded-For":["114.10.5.46"],"X-Forwarded-Proto":["https"],"X-Forwarded-Host":["dms-api.indomobil.co.id"],"X-Forwarded-Port":["443"],"X-NginX-Proxy":["true"]}`, // @p9
		"default",                 // @p10
		"default",                 // @p11
		"default",                 // @p12
		10,                        // @p13
		1,                         // @p14
		"Success",                 // @p15
		"2024-04-03 08:45:28.223", // @p16
		"AdminDMSMobile",          // @p17
	).Row()
	return logSysNo
}
func executeLogbookInsert(db *gorm.DB, params entities.LogbookInsertParams) (int64, error) {
	var output entities.LogbookInsertParams

	query := `
		EXEC api.uspg_logbook_insert
			@LOG_SYS_NO = @p1 OUTPUT,
			@RECORD_STATUS = @p2,
			@API_NAME = @p3,
			@API_METHOD = @p4,
			@API_URL = @p5,
			@API_PARAMETERS = @p6,
			@API_START = @p7,
			@API_FINISH = @p8,
			@API_STATUS = @p9,
			@REQUEST_HEADER = @p10,
			@REQUEST_BODY = @p11,
			@RESPONSE_HEADER = @p12,
			@RESPONSE_BODY = @p13,
			@ROWS_AFFECTED = @p14,
			@SUCCESS = @p15,
			@MESSAGE = @p16,
			@PROCESS_FROM = @p17,
			@PROCESS_BY = @p18
	`

	result := db.Raw(query,
		sql.Out{Dest: &output.LogSysNo}, // @p1 - Output parameter
		params.RecordStatus,             // @p2
		params.ApiName,                  // @p3
		params.ApiMethod,                // @p4
		params.ApiUrl,                   // @p5
		params.ApiParameters,            // @p6
		params.ApiStart,                 // @p7
		params.ApiFinish,                // @p8
		params.ApiStatus,                // @p9
		params.RequestHeader,            // @p10
		params.RequestBody,              // @p11
		params.ResponseHeader,           // @p12
		params.ResponseBody,             // @p13
		params.RowsAffected,             // @p14
		params.Success,                  // @p15
		params.Message,                  // @p16
		params.ProcessFrom,              // @p17
		params.ProcessBy,                // @p18
	).Row()
	if err := result.Err(); err != nil {
		panic(err)
	}

	return output.LogSysNo, nil
}
