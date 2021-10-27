// Code Generated by gadget/xsdk, DO NOT EDIT

package dns

type ListLinesRequest struct {
	Hierarchy  *string `form:"-" json:"-"`
	PageNumber *string `form:"-" json:"-"`
	PageSize   *string `form:"-" json:"-"`
}

type ListLinesResponse struct {
	Lines      []TopLineResponse `form:"Lines" json:"Lines,omitempty"`
	PageNumber *int64            `form:"PageNumber" json:"PageNumber,omitempty"`
	PageSize   *int64            `form:"PageSize" json:"PageSize,omitempty"`
	TotalCount *int64            `form:"TotalCount" json:"TotalCount,omitempty"`
}

type ListRecordAttributesRequest struct {
	ZID *string `form:"-" json:"-"`
}

type ListRecordAttributesResponse struct {
	TTLs  []int64  `form:"TTLs" json:"TTLs,omitempty"`
	Types []string `form:"Types" json:"Types,omitempty"`
}

type CreateRecordRequest struct {
	Host   *string `form:"Host" json:"Host,omitempty"`
	Line   *string `form:"Line" json:"Line,omitempty"`
	TTL    *int64  `form:"TTL" json:"TTL,omitempty"`
	Type   *string `form:"Type" json:"Type,omitempty"`
	Value  *string `form:"Value" json:"Value,omitempty"`
	Weight *int64  `form:"Weight" json:"Weight,omitempty"`
	ZID    *int64  `form:"ZID" json:"ZID,omitempty"`
}

type CreateRecordResponse struct {
	CreatedAt   *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	Enable      *bool    `form:"Enable" json:"Enable,omitempty"`
	FQDN        *string  `form:"FQDN" json:"FQDN,omitempty"`
	Host        *string  `form:"Host" json:"Host,omitempty"`
	Line        *string  `form:"Line" json:"Line,omitempty"`
	Operators   []string `form:"Operators" json:"Operators,omitempty"`
	RecordID    *string  `form:"RecordID" json:"RecordID,omitempty"`
	RecordSetID *string  `form:"RecordSetID" json:"RecordSetID,omitempty"`
	TTL         *int64   `form:"TTL" json:"TTL,omitempty"`
	Type        *string  `form:"Type" json:"Type,omitempty"`
	UpdatedAt   *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	Value       *string  `form:"Value" json:"Value,omitempty"`
	Weight      *int64   `form:"Weight" json:"Weight,omitempty"`
}

type DeleteRecordRequest struct {
	RecordID *string `form:"RecordID" json:"RecordID,omitempty"`
}

type ListRecordSetsRequest struct {
	Host        *string `form:"-" json:"-"`
	PageNumber  *string `form:"-" json:"-"`
	PageSize    *string `form:"-" json:"-"`
	RecordSetID *string `form:"-" json:"-"`
	SearchMode  *string `form:"-" json:"-"`
	ZID         *string `form:"-" json:"-"`
}

type ListRecordSetsResponse struct {
	PageNumber *int64             `form:"PageNumber" json:"PageNumber,omitempty"`
	PageSize   *int64             `form:"PageSize" json:"PageSize,omitempty"`
	RecordSets []TopRecordSetResp `form:"RecordSets" json:"RecordSets,omitempty"`
	TotalCount *int64             `form:"TotalCount" json:"TotalCount,omitempty"`
}

type ListRecordsRequest struct {
	Host        *string `form:"-" json:"-"`
	Line        *string `form:"-" json:"-"`
	PageNumber  *string `form:"-" json:"-"`
	PageSize    *string `form:"-" json:"-"`
	RecordSetID *string `form:"-" json:"-"`
	SearchMode  *string `form:"-" json:"-"`
	SearchOrder *string `form:"-" json:"-"`
	Type        *string `form:"-" json:"-"`
	Value       *string `form:"-" json:"-"`
	ZID         *string `form:"-" json:"-"`
}

type ListRecordsResponse struct {
	PageNumber *int64              `form:"PageNumber" json:"PageNumber,omitempty"`
	PageSize   *int64              `form:"PageSize" json:"PageSize,omitempty"`
	Records    []TopRecordResponse `form:"Records" json:"Records,omitempty"`
	TotalCount *int64              `form:"TotalCount" json:"TotalCount,omitempty"`
}

type QueryRecordRequest struct {
	Host     *string `form:"-" json:"-"`
	Line     *string `form:"-" json:"-"`
	RecordID *string `form:"-" json:"-"`
	Type     *string `form:"-" json:"-"`
	Value    *string `form:"-" json:"-"`
}

type QueryRecordResponse struct {
	CreatedAt   *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	Enable      *bool    `form:"Enable" json:"Enable,omitempty"`
	FQDN        *string  `form:"FQDN" json:"FQDN,omitempty"`
	Host        *string  `form:"Host" json:"Host,omitempty"`
	Line        *string  `form:"Line" json:"Line,omitempty"`
	Operators   []string `form:"Operators" json:"Operators,omitempty"`
	RecordID    *string  `form:"RecordID" json:"RecordID,omitempty"`
	RecordSetID *string  `form:"RecordSetID" json:"RecordSetID,omitempty"`
	TTL         *int64   `form:"TTL" json:"TTL,omitempty"`
	Type        *string  `form:"Type" json:"Type,omitempty"`
	UpdatedAt   *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	Value       *string  `form:"Value" json:"Value,omitempty"`
	Weight      *int64   `form:"Weight" json:"Weight,omitempty"`
}

type SyncFullRecordsRequest struct {
	Await              *bool                                                                `form:"Await" json:"Await,omitempty"`
	DomainWeightEnable *bool                                                                `form:"DomainWeightEnable" json:"DomainWeightEnable,omitempty"`
	IncludeTypes       *string                                                              `form:"IncludeTypes" json:"IncludeTypes,omitempty"`
	SyncConf           *string                                                              `form:"SyncConf" json:"SyncConf,omitempty"`
	Zones              map[string]map[string]map[string]map[string]map[string]AddressConfig `form:"Zones" json:"Zones,omitempty"`
}

type SyncFullRecordsResponse struct {
	TaskID      *string `form:"TaskID" json:"TaskID,omitempty"`
	TotalRecord *int64  `form:"TotalRecord" json:"TotalRecord,omitempty"`
}

type UpdateRecordRequest struct {
	Host     string  `form:"Host" json:"Host"`
	Line     string  `form:"Line" json:"Line"`
	RecordID string  `form:"RecordID" json:"RecordID"`
	TTL      *int64  `form:"TTL" json:"TTL,omitempty"`
	Type     *string `form:"Type" json:"Type,omitempty"`
	Value    *string `form:"Value" json:"Value,omitempty"`
	Weight   *int64  `form:"Weight" json:"Weight,omitempty"`
}

type UpdateRecordResponse struct {
	CreatedAt   *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	Enable      *bool    `form:"Enable" json:"Enable,omitempty"`
	FQDN        *string  `form:"FQDN" json:"FQDN,omitempty"`
	Host        *string  `form:"Host" json:"Host,omitempty"`
	Line        *string  `form:"Line" json:"Line,omitempty"`
	Operators   []string `form:"Operators" json:"Operators,omitempty"`
	RecordID    *string  `form:"RecordID" json:"RecordID,omitempty"`
	RecordSetID *string  `form:"RecordSetID" json:"RecordSetID,omitempty"`
	TTL         *int64   `form:"TTL" json:"TTL,omitempty"`
	Type        *string  `form:"Type" json:"Type,omitempty"`
	UpdatedAt   *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	Value       *string  `form:"Value" json:"Value,omitempty"`
	Weight      *int64   `form:"Weight" json:"Weight,omitempty"`
}

type UpdateRecordSetRequest struct {
	ID            string `form:"ID" json:"ID"`
	WeightEnabled bool   `form:"WeightEnabled" json:"WeightEnabled"`
}

type UpdateRecordSetResponse struct {
	FQDN          *string `form:"FQDN" json:"FQDN,omitempty"`
	Host          *string `form:"Host" json:"Host,omitempty"`
	ID            *string `form:"ID" json:"ID,omitempty"`
	Line          *string `form:"Line" json:"Line,omitempty"`
	Type          *string `form:"Type" json:"Type,omitempty"`
	WeightEnabled *bool   `form:"WeightEnabled" json:"WeightEnabled,omitempty"`
}

type UpdateRecordStatusRequest struct {
	Enable   *bool   `form:"Enable" json:"Enable,omitempty"`
	RecordID *string `form:"RecordID" json:"RecordID,omitempty"`
}

type UpdateRecordStatusResponse struct {
	CreatedAt   *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	Enable      *bool    `form:"Enable" json:"Enable,omitempty"`
	FQDN        *string  `form:"FQDN" json:"FQDN,omitempty"`
	Host        *string  `form:"Host" json:"Host,omitempty"`
	Line        *string  `form:"Line" json:"Line,omitempty"`
	Operators   []string `form:"Operators" json:"Operators,omitempty"`
	RecordID    *string  `form:"RecordID" json:"RecordID,omitempty"`
	RecordSetID *string  `form:"RecordSetID" json:"RecordSetID,omitempty"`
	TTL         *int64   `form:"TTL" json:"TTL,omitempty"`
	Type        *string  `form:"Type" json:"Type,omitempty"`
	UpdatedAt   *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	Value       *string  `form:"Value" json:"Value,omitempty"`
	Weight      *int64   `form:"Weight" json:"Weight,omitempty"`
}

type ListDomainStatisticsRequest struct {
	End        *string `form:"-" json:"-"`
	Name       *string `form:"-" json:"-"`
	PageNumber *string `form:"-" json:"-"`
	PageSize   *string `form:"-" json:"-"`
	SearchMode *string `form:"-" json:"-"`
	Start      *string `form:"-" json:"-"`
	Threshold  *string `form:"-" json:"-"`
	ZID        *string `form:"-" json:"-"`
}

type ListDomainStatisticsResponse struct {
	Data  []GroupStat `form:"data" json:"data,omitempty"`
	Page  *int64      `form:"page" json:"page,omitempty"`
	Size  *int64      `form:"size" json:"size,omitempty"`
	Total *int64      `form:"total" json:"total,omitempty"`
}

type ListZoneStatisticsRequest struct {
	End        *string `form:"-" json:"-"`
	Name       *string `form:"-" json:"-"`
	PageNumber *string `form:"-" json:"-"`
	PageSize   *string `form:"-" json:"-"`
	SearchMode *string `form:"-" json:"-"`
	Start      *string `form:"-" json:"-"`
	Threshold  *string `form:"-" json:"-"`
}

type ListZoneStatisticsResponse struct {
	Data  []GroupStat `form:"data" json:"data,omitempty"`
	Page  *int64      `form:"page" json:"page,omitempty"`
	Size  *int64      `form:"size" json:"size,omitempty"`
	Total *int64      `form:"total" json:"total,omitempty"`
}

type QueryDomainStatisticsRequest struct {
	End   *string `form:"-" json:"-"`
	Name  *string `form:"-" json:"-"`
	Start *string `form:"-" json:"-"`
	ZID   *string `form:"-" json:"-"`
}

type QueryDomainStatisticsResponse []Stat

type QueryZoneStatisticsRequest struct {
	End   *string `form:"-" json:"-"`
	Start *string `form:"-" json:"-"`
	ZID   *string `form:"-" json:"-"`
}

type QueryZoneStatisticsResponse []Stat

type CalculatePriceRequest struct {
	CouponID   *string `form:"-" json:"-"`
	InstanceID *string `form:"-" json:"-"`
	OrderType  *string `form:"-" json:"-"`
	Times      *string `form:"-" json:"-"`
}

type CalculatePriceResponse struct {
	PayableAmount *int64 `form:"PayableAmount" json:"PayableAmount,omitempty"`
}

type CreateNewPreorderRequest struct {
	Times *int64 `form:"Times" json:"Times,omitempty"`
	ZID   *int64 `form:"ZID" json:"ZID,omitempty"`
}

type CreateNewPreorderResponse struct {
	PreOrderNumber *string `form:"PreOrderNumber" json:"PreOrderNumber,omitempty"`
}

type CreateRenewPreorderRequest struct {
	InstanceID *string `form:"InstanceID" json:"InstanceID,omitempty"`
	Times      *int64  `form:"Times" json:"Times,omitempty"`
}

type CreateRenewPreorderResponse struct {
	PreOrderNumber *string `form:"PreOrderNumber" json:"PreOrderNumber,omitempty"`
}

type CreateTerminatePreorderRequest struct {
	InstanceID *string `form:"InstanceID" json:"InstanceID,omitempty"`
}

type CreateTerminatePreorderResponse struct {
	PreOrderNumber *string `form:"PreOrderNumber" json:"PreOrderNumber,omitempty"`
}

type CreateZoneRequest struct {
	Remark   *string `form:"Remark" json:"Remark,omitempty"`
	ZoneName string  `form:"ZoneName" json:"ZoneName"`
}

type CreateZoneResponse struct {
	ConfigurationCode *string `form:"ConfigurationCode" json:"ConfigurationCode,omitempty"`
	CreatedAt         *string `form:"CreatedAt" json:"CreatedAt,omitempty"`
	ExpiredTime       *int64  `form:"ExpiredTime" json:"ExpiredTime,omitempty"`
	InstanceID        *string `form:"InstanceID" json:"InstanceID,omitempty"`
	LastOperator      *string `form:"LastOperator" json:"LastOperator,omitempty"`
	RecordCount       *int64  `form:"RecordCount" json:"RecordCount,omitempty"`
	Remark            *string `form:"Remark" json:"Remark,omitempty"`
	UpdatedAt         *string `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	ZID               *int64  `form:"ZID" json:"ZID,omitempty"`
	ZoneName          *string `form:"ZoneName" json:"ZoneName,omitempty"`
}

type DeleteZoneRequest struct {
	ZID *int64 `form:"ZID" json:"ZID,omitempty"`
}

type ListZonesRequest struct {
	Key        *string `form:"-" json:"-"`
	PageNumber *string `form:"-" json:"-"`
	PageSize   *string `form:"-" json:"-"`
	SearchMode *string `form:"-" json:"-"`
}

type ListZonesResponse struct {
	Total *int64            `form:"Total" json:"Total,omitempty"`
	Zones []TopZoneResponse `form:"Zones" json:"Zones,omitempty"`
}

type QueryZoneRequest struct {
	ZIDrequired *string `form:"-" json:"-"`
}

type QueryZoneResponse struct {
	ConfigurationCode *string `form:"ConfigurationCode" json:"ConfigurationCode,omitempty"`
	ExpiredTime       *int64  `form:"ExpiredTime" json:"ExpiredTime,omitempty"`
	InstanceNo        *string `form:"InstanceNo" json:"InstanceNo,omitempty"`
	Remark            *string `form:"Remark" json:"Remark,omitempty"`
	Status            *int64  `form:"Status" json:"Status,omitempty"`
	ZoneName          *string `form:"ZoneName" json:"ZoneName,omitempty"`
}

type UpdateZoneRequest struct {
	Remark *string `form:"Remark" json:"Remark,omitempty"`
	ZID    *int64  `form:"ZID" json:"ZID,omitempty"`
}

type UpdateZoneResponse struct {
	ConfigurationCode *string `form:"ConfigurationCode" json:"ConfigurationCode,omitempty"`
	CreatedAt         *string `form:"CreatedAt" json:"CreatedAt,omitempty"`
	ExpiredTime       *int64  `form:"ExpiredTime" json:"ExpiredTime,omitempty"`
	InstanceID        *string `form:"InstanceID" json:"InstanceID,omitempty"`
	LastOperator      *string `form:"LastOperator" json:"LastOperator,omitempty"`
	RecordCount       *int64  `form:"RecordCount" json:"RecordCount,omitempty"`
	Remark            *string `form:"Remark" json:"Remark,omitempty"`
	UpdatedAt         *string `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	ZID               *int64  `form:"ZID" json:"ZID,omitempty"`
	ZoneName          *string `form:"ZoneName" json:"ZoneName,omitempty"`
}

type TopLineResponse struct {
	Children    []TopLineResponse `form:"Children" json:"Children,omitempty"`
	FatherValue *string           `form:"FatherValue" json:"FatherValue,omitempty"`
	Name        *string           `form:"Name" json:"Name,omitempty"`
	Value       *string           `form:"Value" json:"Value,omitempty"`
}

type TopRecordSetResp struct {
	FQDN          *string `form:"FQDN" json:"FQDN,omitempty"`
	Host          *string `form:"Host" json:"Host,omitempty"`
	ID            *string `form:"ID" json:"ID,omitempty"`
	Line          *string `form:"Line" json:"Line,omitempty"`
	Type          *string `form:"Type" json:"Type,omitempty"`
	WeightEnabled *bool   `form:"WeightEnabled" json:"WeightEnabled,omitempty"`
}

type TopRecordResponse struct {
	CreatedAt   *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	Enable      *bool    `form:"Enable" json:"Enable,omitempty"`
	FQDN        *string  `form:"FQDN" json:"FQDN,omitempty"`
	Host        *string  `form:"Host" json:"Host,omitempty"`
	Line        *string  `form:"Line" json:"Line,omitempty"`
	Operators   []string `form:"Operators" json:"Operators,omitempty"`
	RecordID    *string  `form:"RecordID" json:"RecordID,omitempty"`
	RecordSetID *string  `form:"RecordSetID" json:"RecordSetID,omitempty"`
	TTL         *int64   `form:"TTL" json:"TTL,omitempty"`
	Type        *string  `form:"Type" json:"Type,omitempty"`
	UpdatedAt   *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	Value       *string  `form:"Value" json:"Value,omitempty"`
	Weight      *int64   `form:"Weight" json:"Weight,omitempty"`
}

type AddressConfig struct {
	Disable *bool  `form:"disable" json:"disable,omitempty"`
	TTL     *int64 `form:"ttl" json:"ttl,omitempty"`
	Weight  *int64 `form:"weight" json:"weight,omitempty"`
}

type GroupStat struct {
	ZID       *int64  `form:"ZID" json:"ZID,omitempty"`
	Error     *int64  `form:"error" json:"error,omitempty"`
	Name      *string `form:"name" json:"name,omitempty"`
	NotExists *int64  `form:"not_exists" json:"not_exists,omitempty"`
	Success   *int64  `form:"success" json:"success,omitempty"`
	Timestamp *int64  `form:"timestamp" json:"timestamp,omitempty"`
	Total     *int64  `form:"total" json:"total,omitempty"`
}

type Stat struct {
	Error     *int64 `form:"error" json:"error,omitempty"`
	NotExists *int64 `form:"not_exists" json:"not_exists,omitempty"`
	Success   *int64 `form:"success" json:"success,omitempty"`
	Timestamp *int64 `form:"timestamp" json:"timestamp,omitempty"`
	Total     *int64 `form:"total" json:"total,omitempty"`
}

type TopZoneResponse struct {
	ConfigurationCode *string `form:"ConfigurationCode" json:"ConfigurationCode,omitempty"`
	CreatedAt         *string `form:"CreatedAt" json:"CreatedAt,omitempty"`
	ExpiredTime       *int64  `form:"ExpiredTime" json:"ExpiredTime,omitempty"`
	InstanceID        *string `form:"InstanceID" json:"InstanceID,omitempty"`
	LastOperator      *string `form:"LastOperator" json:"LastOperator,omitempty"`
	RecordCount       *int64  `form:"RecordCount" json:"RecordCount,omitempty"`
	Remark            *string `form:"Remark" json:"Remark,omitempty"`
	UpdatedAt         *string `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	ZID               *int64  `form:"ZID" json:"ZID,omitempty"`
	ZoneName          *string `form:"ZoneName" json:"ZoneName,omitempty"`
}