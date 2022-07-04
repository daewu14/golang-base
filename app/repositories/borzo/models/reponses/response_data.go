package reponses

import "time"

type ResponseDataShowOrder struct {
	IsSuccessful bool `json:"is_successful"`
	Orders       []struct {
		Type                               string      `json:"type"`
		OrderID                            int         `json:"order_id"`
		OrderName                          string      `json:"order_name"`
		VehicleTypeID                      int         `json:"vehicle_type_id"`
		CreatedDatetime                    time.Time   `json:"created_datetime"`
		FinishDatetime                     interface{} `json:"finish_datetime"`
		Status                             string      `json:"status"`
		StatusDescription                  string      `json:"status_description"`
		Matter                             string      `json:"matter"`
		TotalWeightKg                      int         `json:"total_weight_kg"`
		IsClientNotificationEnabled        bool        `json:"is_client_notification_enabled"`
		IsContactPersonNotificationEnabled bool        `json:"is_contact_person_notification_enabled"`
		LoadersCount                       int         `json:"loaders_count"`
		BackpaymentDetails                 interface{} `json:"backpayment_details"`
		Points                             []struct {
			PointType                string      `json:"point_type"`
			PointID                  int         `json:"point_id"`
			DeliveryID               interface{} `json:"delivery_id"`
			ClientOrderID            interface{} `json:"client_order_id"`
			Address                  string      `json:"address"`
			Latitude                 string      `json:"latitude"`
			Longitude                string      `json:"longitude"`
			RequiredStartDatetime    time.Time   `json:"required_start_datetime"`
			RequiredFinishDatetime   time.Time   `json:"required_finish_datetime"`
			ArrivalStartDatetime     interface{} `json:"arrival_start_datetime"`
			ArrivalFinishDatetime    interface{} `json:"arrival_finish_datetime"`
			EstimatedArrivalDatetime interface{} `json:"estimated_arrival_datetime"`
			CourierVisitDatetime     interface{} `json:"courier_visit_datetime"`
			ContactPerson            struct {
				Name  string `json:"name"`
				Phone string `json:"phone"`
			} `json:"contact_person"`
			TakingAmount                        string        `json:"taking_amount"`
			BuyoutAmount                        string        `json:"buyout_amount"`
			Note                                interface{}   `json:"note"`
			PreviousPointDrivingDistanceMeters  int           `json:"previous_point_driving_distance_meters"`
			Packages                            []interface{} `json:"packages"`
			IsCodCashVoucherRequired            bool          `json:"is_cod_cash_voucher_required"`
			PlacePhotoURL                       interface{}   `json:"place_photo_url"`
			SignPhotoURL                        interface{}   `json:"sign_photo_url"`
			TrackingURL                         string        `json:"tracking_url"`
			Checkin                             interface{}   `json:"checkin"`
			IsReturnPoint                       bool          `json:"is_return_point"`
			IsOrderPaymentHere                  bool          `json:"is_order_payment_here"`
			BuildingNumber                      interface{}   `json:"building_number"`
			EntranceNumber                      interface{}   `json:"entrance_number"`
			IntercomCode                        interface{}   `json:"intercom_code"`
			FloorNumber                         interface{}   `json:"floor_number"`
			ApartmentNumber                     interface{}   `json:"apartment_number"`
			InvisibleMileNavigationInstructions interface{}   `json:"invisible_mile_navigation_instructions"`
		} `json:"points"`
		PaymentAmount              string      `json:"payment_amount"`
		DeliveryFeeAmount          string      `json:"delivery_fee_amount"`
		WeightFeeAmount            string      `json:"weight_fee_amount"`
		InsuranceAmount            string      `json:"insurance_amount"`
		InsuranceFeeAmount         string      `json:"insurance_fee_amount"`
		LoadingFeeAmount           string      `json:"loading_fee_amount"`
		MoneyTransferFeeAmount     string      `json:"money_transfer_fee_amount"`
		OvernightFeeAmount         string      `json:"overnight_fee_amount"`
		DoorToDoorFeeAmount        string      `json:"door_to_door_fee_amount"`
		PromoCodeDiscountAmount    string      `json:"promo_code_discount_amount"`
		BackpaymentAmount          string      `json:"backpayment_amount"`
		CodFeeAmount               string      `json:"cod_fee_amount"`
		BackpaymentPhotoURL        interface{} `json:"backpayment_photo_url"`
		ItineraryDocumentURL       interface{} `json:"itinerary_document_url"`
		WaybillDocumentURL         interface{} `json:"waybill_document_url"`
		Courier                    interface{} `json:"courier"`
		IsMotoboxRequired          bool        `json:"is_motobox_required"`
		PaymentMethod              string      `json:"payment_method"`
		BankCardID                 interface{} `json:"bank_card_id"`
		AppliedPromoCode           interface{} `json:"applied_promo_code"`
		IsReturnRequired           bool        `json:"is_return_required"`
		IntercityDeliveryFeeAmount string      `json:"intercity_delivery_fee_amount"`
		SuburbanDeliveryFeeAmount  string      `json:"suburban_delivery_fee_amount"`
		DiscountAmount             string      `json:"discount_amount"`
	} `json:"orders"`
	OrdersCount int `json:"orders_count"`
}
