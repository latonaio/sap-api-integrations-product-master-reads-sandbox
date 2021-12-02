package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-master-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToProduct(raw []byte, l *logger.Logger) (*Product, error) {
	pm := &responses.Product{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Product. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Product{
		Product:             data.Product,
		ProductDescription:  data.ToDescription.Deferred.URI,
		BaseUnit:            data.BaseUnit,
		ValidityStartDate:   data.ValidityStartDate,
		ProductGroup:        data.ProductGroup,
		Division:            data.Division,
		GrossWeight:         data.GrossWeight,
		WeightUnit:          data.WeightUnit,
		SizeOrDimensionText: data.SizeOrDimensionText,
		ProductStandardID:   data.ProductStandardID,
	}, nil
}

func ConvertToProductDesc(raw []byte, l *logger.Logger) (*ProductDesc, error) {
	desc := &responses.ProductDesc{}
	err := json.Unmarshal(raw, desc)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductDesc. unmarshal error: %w", err)
	}

	return &ProductDesc{
		Product:            desc.D.Product,
		ProductDescription: desc.D.ProductDescription,
	}, nil
}

func ConvertToPlant(raw []byte, l *logger.Logger) (*Plant, error) {
	pd := &responses.Plant{}
	err := json.Unmarshal(raw, pd)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Plant. unmarshal error: %#v", err)
	}
	if len(pd.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pd.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pd.D.Results))
	}
	data := pd.D.Results[0]

	return &Plant{
		Product:                       data.Product,
		Plant:                         data.Plant,
		PurchasingGroup:               data.PurchasingGroup,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		AvailabilityCheckType:         data.AvailabilityCheckType,
		ProfitCenter:                  data.ProfitCenter,
		GoodsReceiptDuration:          data.GoodsReceiptDuration,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		IsBatchManagementRequired:     data.IsBatchManagementRequired,
		ProcurementType:               data.ProcurementType,
		IsInternalBatchManaged:        data.IsInternalBatchManaged,
		GoodsIssueUnit:                data.GoodsIssueUnit,
		MaterialFreightGroup:          data.MaterialFreightGroup,
		ProductLogisticsHandlingGroup: data.ProductLogisticsHandlingGroup,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}, nil
}

func ConvertToMRPArea(raw []byte, l *logger.Logger) (*MRPArea, error) {
	ma := &responses.MRPArea{}
	err := json.Unmarshal(raw, ma)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MRPArea. unmarshal error: %w", err)
	}
	if len(ma.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(ma.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(ma.D.Results))
	}
	data := ma.D.Results[0]

	return &MRPArea{
		Product:                       data.Product,
		Plant:                         data.Plant,
		MRPArea:                       data.MRPArea,
		MRPType:                       data.MRPType,
		MRPResponsible:                data.MRPResponsible,
		MRPGroup:                      data.MRPGroup,
		ReorderThresholdQuantity:      data.ReorderThresholdQuantity,
		PlanningTimeFence:             data.PlanningTimeFence,
		LotSizingProcedure:            data.LotSizingProcedure,
		LotSizeRoundingQuantity:       data.LotSizeRoundingQuantity,
		MinimumLotSizeQuantity:        data.MinimumLotSizeQuantity,
		MaximumLotSizeQuantity:        data.MaximumLotSizeQuantity,
		MaximumStockQuantity:          data.MaximumStockQuantity,
		ProcurementSubType:            data.ProcurementSubType,
		DfltStorageLocationExtProcmt:  data.DfltStorageLocationExtProcmt,
		MRPPlanningCalendar:           data.MRPPlanningCalendar,
		SafetyStockQuantity:           data.SafetyStockQuantity,
		SafetyDuration:                data.SafetyDuration,
		FixedLotSizeQuantity:          data.FixedLotSizeQuantity,
		PlannedDeliveryDurationInDays: data.PlannedDeliveryDurationInDays,
		StorageLocation:               data.StorageLocation,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}, nil
}

func ConvertToProcurement(raw []byte, l *logger.Logger) (*Procurement, error) {
	pm := &responses.Procurement{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Procurement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Procurement{
		Product:                     data.Product,
		Plant:                       data.Plant,
		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:        data.IsSourceListRequired,
	}, nil
}

func ConvertToWorkScheduling(raw []byte, l *logger.Logger) (*WorkScheduling, error) {
	ws := &responses.WorkScheduling{}
	err := json.Unmarshal(raw, ws)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to WorkScheduling. unmarshal error: %w", err)
	}
	if len(ws.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(ws.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(ws.D.Results))
	}
	data := ws.D.Results[0]

	return &WorkScheduling{
		Product:                       data.Product,
		Plant:                         data.Plant,
		ProductionInvtryManagedLoc:    data.ProductionInvtryManagedLoc,
		ProductProcessingTime:         data.ProductProcessingTime,
		ProductionSupervisor:          data.ProductionSupervisor,
		ProductProductionQuantityUnit: data.ProductProductionQuantityUnit,
		ProdnOrderIsBatchRequired:     data.ProdnOrderIsBatchRequired,
		MatlCompIsMarkedForBackflush:  data.MatlCompIsMarkedForBackflush,
		ProductionSchedulingProfile:   data.ProductionSchedulingProfile,
	}, nil
}

func ConvertToSalesPlant(raw []byte, l *logger.Logger) (*SalesPlant, error) {
	pps := &responses.SalesPlant{}
	err := json.Unmarshal(raw, pps)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesPlant. unmarshal error: %w", err)
	}
	if len(pps.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pps.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pps.D.Results))
	}
	data := pps.D.Results[0]

	return &SalesPlant{
		Product:               data.Product,
		Plant:                 data.Plant,
		LoadingGroup:          data.LoadingGroup,
		AvailabilityCheckType: data.AvailabilityCheckType,
	}, nil
}

func ConvertToProcurment(raw []byte, l *logger.Logger) (*Procurement, error) {
	pm := &responses.Procurement{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Procurement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Procurement{
		Product:                     data.Product,
		Plant:                       data.Plant,
		IsAutoPurOrdCreationAllowed: data.IsAutoPurOrdCreationAllowed,
		IsSourceListRequired:        data.IsSourceListRequired,
	}, nil
}

func ConvertToAccounting(raw []byte, l *logger.Logger) (*Accounting, error) {
	pad := &responses.Accounting{}
	err := json.Unmarshal(raw, pad)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Accounting. unmarshal error: %w", err)
	}
	if len(pad.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pad.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pad.D.Results))
	}
	data := pad.D.Results[0]

	return &Accounting{
		Product:             data.Product,
		ValuationArea:       data.ValuationArea,
		ValuationClass:      data.ValuationClass,
		StandardPrice:       data.StandardPrice,
		PriceUnitQty:        data.PriceUnitQty,
		MovingAveragePrice:  data.MovingAveragePrice,
		PriceLastChangeDate: data.PriceLastChangeDate,
		PlannedPrice:        data.PlannedPrice,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}, nil
}

func ConvertToSalesOrganization(raw []byte, l *logger.Logger) (*SalesOrganization, error) {
	psd := &responses.SalesOrganization{}
	err := json.Unmarshal(raw, psd)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to SalesOrganization. unmarshal error: %w", err)
	}
	if len(psd.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(psd.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(psd.D.Results))
	}
	data := psd.D.Results[0]

	return &SalesOrganization{
		Product:                        data.Product,
		ProductSalesOrg:                data.ProductSalesOrg,
		ProductDistributionChnl:        data.ProductDistributionChnl,
		SupplyingPlant:                 data.SupplyingPlant,
		PriceSpecificationProductGroup: data.PriceSpecificationProductGroup,
		AccountDetnProductGroup:        data.AccountDetnProductGroup,
		ItemCategoryGroup:              data.ItemCategoryGroup,
		SalesMeasureUnit:               data.SalesMeasureUnit,
		ProductHierarchy:               data.ProductHierarchy,
		IsMarkedForDeletion:            data.IsMarkedForDeletion,
	}, nil
}
