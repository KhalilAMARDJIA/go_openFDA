package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://api.fda.gov/device/event.json?search="

type openFDA_event struct {
	Meta struct {
		Disclaimer  string `json:"disclaimer"`
		Terms       string `json:"terms"`
		License     string `json:"license"`
		LastUpdated string `json:"last_updated"`
		Results     struct {
			Skip  int `json:"skip"`
			Limit int `json:"limit"`
			Total int `json:"total"`
		} `json:"results"`
	} `json:"meta"`
	Results []struct {
		ManufacturerContactZipExt   string   `json:"manufacturer_contact_zip_ext"`
		ManufacturerG1Address2      string   `json:"manufacturer_g1_address_2"`
		EventLocation               string   `json:"event_location"`
		ReportToFda                 string   `json:"report_to_fda"`
		ManufacturerContactTName    string   `json:"manufacturer_contact_t_name"`
		ManufacturerContactState    string   `json:"manufacturer_contact_state"`
		ManufacturerLinkFlag        string   `json:"manufacturer_link_flag"`
		ManufacturerContactAddress2 string   `json:"manufacturer_contact_address_2"`
		ManufacturerG1City          string   `json:"manufacturer_g1_city"`
		ManufacturerContactAddress1 string   `json:"manufacturer_contact_address_1"`
		ManufacturerContactPcity    string   `json:"manufacturer_contact_pcity"`
		EventType                   string   `json:"event_type"`
		ReportNumber                string   `json:"report_number"`
		TypeOfReport                []string `json:"type_of_report"`
		ProductProblemFlag          string   `json:"product_problem_flag"`
		DateReceived                string   `json:"date_received"`
		ManufacturerAddress2        string   `json:"manufacturer_address_2"`
		PmaPmnNumber                string   `json:"pma_pmn_number"`
		ReprocessedAndReusedFlag    string   `json:"reprocessed_and_reused_flag"`
		ManufacturerAddress1        string   `json:"manufacturer_address_1"`
		ExemptionNumber             string   `json:"exemption_number"`
		ManufacturerContactZipCode  string   `json:"manufacturer_contact_zip_code"`
		ReporterOccupationCode      string   `json:"reporter_occupation_code"`
		ManufacturerContactPlocal   string   `json:"manufacturer_contact_plocal"`
		ManufacturerContactLName    string   `json:"manufacturer_contact_l_name"`
		SourceType                  []string `json:"source_type"`
		DistributorZipCodeExt       string   `json:"distributor_zip_code_ext"`
		ManufacturerG1PostalCode    string   `json:"manufacturer_g1_postal_code"`
		DateFacilityAware           string   `json:"date_facility_aware"`
		ManufacturerG1State         string   `json:"manufacturer_g1_state"`
		ReporterCountryCode         string   `json:"reporter_country_code"`
		ManufacturerContactAreaCode string   `json:"manufacturer_contact_area_code"`
		DateAdded                   string   `json:"date_added"`
		ManufacturerContactFName    string   `json:"manufacturer_contact_f_name"`
		PreviousUseCode             string   `json:"previous_use_code"`
		Device                      []struct {
			DeviceEventKey                string `json:"device_event_key"`
			ImplantFlag                   string `json:"implant_flag"`
			DateRemovedFlag               string `json:"date_removed_flag"`
			DeviceSequenceNumber          string `json:"device_sequence_number"`
			DateReceived                  string `json:"date_received"`
			BrandName                     string `json:"brand_name"`
			GenericName                   string `json:"generic_name"`
			ManufacturerDName             string `json:"manufacturer_d_name"`
			ManufacturerDAddress1         string `json:"manufacturer_d_address_1"`
			ManufacturerDAddress2         string `json:"manufacturer_d_address_2"`
			ManufacturerDCity             string `json:"manufacturer_d_city"`
			ManufacturerDState            string `json:"manufacturer_d_state"`
			ManufacturerDZipCode          string `json:"manufacturer_d_zip_code"`
			ManufacturerDZipCodeExt       string `json:"manufacturer_d_zip_code_ext"`
			ManufacturerDCountry          string `json:"manufacturer_d_country"`
			ManufacturerDPostalCode       string `json:"manufacturer_d_postal_code"`
			DeviceOperator                string `json:"device_operator"`
			ModelNumber                   string `json:"model_number"`
			CatalogNumber                 string `json:"catalog_number"`
			LotNumber                     string `json:"lot_number"`
			OtherIDNumber                 string `json:"other_id_number"`
			DeviceAvailability            string `json:"device_availability"`
			DeviceReportProductCode       string `json:"device_report_product_code"`
			DeviceAgeText                 string `json:"device_age_text"`
			DeviceEvaluatedByManufacturer string `json:"device_evaluated_by_manufacturer"`
			CombinationProductFlag        string `json:"combination_product_flag"`
			Openfda                       struct {
				DeviceName                  string `json:"device_name"`
				MedicalSpecialtyDescription string `json:"medical_specialty_description"`
				RegulationNumber            string `json:"regulation_number"`
				DeviceClass                 string `json:"device_class"`
			} `json:"openfda"`
		} `json:"device"`
		ProductProblems                []string `json:"product_problems"`
		ManufacturerZipCode            string   `json:"manufacturer_zip_code"`
		ManufacturerContactCountry     string   `json:"manufacturer_contact_country"`
		DateChanged                    string   `json:"date_changed"`
		HealthProfessional             string   `json:"health_professional"`
		SummaryReportFlag              string   `json:"summary_report_flag"`
		ManufacturerG1ZipCodeExt       string   `json:"manufacturer_g1_zip_code_ext"`
		ManufacturerContactExtension   string   `json:"manufacturer_contact_extension"`
		ManufacturerCity               string   `json:"manufacturer_city"`
		ManufacturerContactPhoneNumber string   `json:"manufacturer_contact_phone_number"`
		Patient                        []struct {
			PatientSequenceNumber   string   `json:"patient_sequence_number"`
			DateReceived            string   `json:"date_received"`
			SequenceNumberTreatment []string `json:"sequence_number_treatment"`
			SequenceNumberOutcome   []string `json:"sequence_number_outcome"`
			PatientProblems         []string `json:"patient_problems"`
		} `json:"patient"`
		DistributorCity          string   `json:"distributor_city"`
		DateReport               string   `json:"date_report"`
		InitialReportToFda       string   `json:"initial_report_to_fda"`
		DistributorState         string   `json:"distributor_state"`
		EventKey                 string   `json:"event_key"`
		ManufacturerG1Country    string   `json:"manufacturer_g1_country"`
		ManufacturerContactCity  string   `json:"manufacturer_contact_city"`
		MdrReportKey             string   `json:"mdr_report_key"`
		RemovalCorrectionNumber  string   `json:"removal_correction_number"`
		NumberDevicesInEvent     string   `json:"number_devices_in_event"`
		DateManufacturerReceived string   `json:"date_manufacturer_received"`
		ManufacturerName         string   `json:"manufacturer_name"`
		ReportSourceCode         string   `json:"report_source_code"`
		RemedialAction           []string `json:"remedial_action"`
		ManufacturerG1ZipCode    string   `json:"manufacturer_g1_zip_code"`
		ReportToManufacturer     string   `json:"report_to_manufacturer"`
		ManufacturerZipCodeExt   string   `json:"manufacturer_zip_code_ext"`
		ManufacturerG1Name       string   `json:"manufacturer_g1_name"`
		AdverseEventFlag         string   `json:"adverse_event_flag"`
		DistributorAddress1      string   `json:"distributor_address_1"`
		ManufacturerState        string   `json:"manufacturer_state"`
		DistributorAddress2      string   `json:"distributor_address_2"`
		ManufacturerPostalCode   string   `json:"manufacturer_postal_code"`
		SingleUseFlag            string   `json:"single_use_flag"`
		ManufacturerCountry      string   `json:"manufacturer_country"`
		MdrText                  []struct {
			MdrTextKey            string `json:"mdr_text_key"`
			TextTypeCode          string `json:"text_type_code"`
			PatientSequenceNumber string `json:"patient_sequence_number"`
			Text                  string `json:"text"`
		} `json:"mdr_text"`
		NumberPatientsInEvent         string `json:"number_patients_in_event"`
		DistributorName               string `json:"distributor_name"`
		ManufacturerG1Address1        string `json:"manufacturer_g1_address_1"`
		DistributorZipCode            string `json:"distributor_zip_code"`
		ManufacturerContactPostalCode string `json:"manufacturer_contact_postal_code"`
		ManufacturerContactExchange   string `json:"manufacturer_contact_exchange"`
		ManufacturerContactPcountry   string `json:"manufacturer_contact_pcountry"`
	} `json:"results"`
}

func main() {
	// query construct
	fmt.Println("Enter openFDA query: ")
	var query string
	fmt.Scanln(&query)
	var full_query string
	full_query = baseURL + query + "&limit=1000"

	// openFDA API request
	response, err := http.Get(full_query)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Write raw data into json file
	ioutil.WriteFile("openFDA_raw_data.json", responseData, os.ModePerm)

	content := openFDA_event{}
	json.Unmarshal([]byte(responseData), &content)
	for i := 0; i < len(content.Results); i++ {
		fmt.Printf("%+v\n", content.Results[i].Device[0].BrandName)
	}

}
