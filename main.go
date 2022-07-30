package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

const baseURL = "https://api.fda.gov/device/event.json?search="
const limit = "1000"

type openFDA_event_schema struct {
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

func query_to_json(query string) []byte {

	// openFDA API request
	response, err := http.Get(query)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

func query_constructed() string {
	// query construct
	app := app.New()
	window := app.NewWindow("openFDA event database search")
	window.Resize(fyne.NewSize(800, 100))
	window.CenterOnScreen()
	input := widget.NewEntry()
	input.SetPlaceHolder("Query")

	content := container.NewVBox(input, widget.NewButton("Search", func() {

		var full_query string = baseURL + input.Text + "&limit=" + limit

		// Write query to file
		f, err := os.Create("Last_query.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(full_query)

		if err2 != nil {
			log.Fatal(err2)
		}
	}))

	window.SetContent(content)
	window.ShowAndRun()
	window.Close()
	return "search"
}

func find_meta() (string, int, int) {
	query_constructed()

	last_query_file, err := os.ReadFile("Last_query.txt") // Read latest query file
	if err != nil {
		fmt.Print(err)
	}

	meta_query := string(last_query_file)
	responseData := query_to_json(meta_query)
	content := openFDA_event_schema{}
	json.Unmarshal([]byte(responseData), &content)

	// Show metadata in console
	fmt.Print("Results found: ", content.Meta.Results.Total, " Last update in: ", content.Meta.LastUpdated, "\n", "Warning only the first 27000th records will be taken into account (openFDA paging limit)")
	limit_int, err := strconv.Atoi(limit) // convert limit string to int
	if err != nil {
		fmt.Println(err)
	}

	if content.Meta.Results.Total <= limit_int {
		var skips_required int = 0
		return meta_query, skips_required, limit_int // skips_required variable
	} else {
		var skips_required int = content.Meta.Results.Total / limit_int
		return meta_query, skips_required, limit_int // skips_required variable

	}
}

func get_data() []openFDA_event_schema {
	meta_query, skips_required, limit_int := find_meta()
	var query_array []string
	var responseData []byte

	for i := 0; i <= skips_required; i++ {
		skip_string := strconv.Itoa(i * limit_int)
		query_per_page := meta_query + "&skip=" + skip_string
		query_array = append(query_array, query_per_page)
	}
	var all_content []openFDA_event_schema
	for _, query_per_page := range query_array {
		responseData = query_to_json(query_per_page)
		content := openFDA_event_schema{}
		json.Unmarshal([]byte(responseData), &content)

		all_content = append(all_content, content)
	}

	return all_content
}

func data_to_csv() {
	data := get_data()

	csvFile, err := os.Create("./output_data/openFDA_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	// Define header row
	headerRow := []string{
		"report_number\tdate_received\tmanufacturer_name\tbrand_name\tpatient_problems\tproduct_problems\ttext\t",
	}
	writer.Write(headerRow)

	for _, data_page := range data {
		for _, usance := range data_page.Results {
			writer.Comma = ';'
			var row []string
			var nest_row []string
			row = append(row, usance.ReportNumber)
			row = append(row, usance.DateReceived)
			row = append(row, usance.Device[0].ManufacturerDName)
			row = append(row, usance.Device[0].BrandName)
			row = append(row, strings.Join(usance.Patient[0].PatientProblems, "|"))
			row = append(row, strings.Join(usance.ProductProblems, "|"))
			for _, txt := range usance.MdrText {
				nest_row = append(nest_row, txt.Text)
			}
			row = append(row, strings.Join(nest_row, "|"))
			writer.Write(row)
			writer.Flush() // Data flush
		}
	}
}

func main() {
	data_to_csv()

}
