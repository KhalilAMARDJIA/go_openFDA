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
)

const baseURL = "https://api.fda.gov/drug/event.json?search="
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
		Safetyreportversion     string `json:"safetyreportversion"`
		Safetyreportid          string `json:"safetyreportid"`
		Primarysourcecountry    string `json:"primarysourcecountry"`
		Occurcountry            string `json:"occurcountry"`
		Transmissiondateformat  string `json:"transmissiondateformat"`
		Transmissiondate        string `json:"transmissiondate"`
		Reporttype              string `json:"reporttype"`
		Serious                 string `json:"serious"`
		Seriousnessother        string `json:"seriousnessother"`
		Receivedateformat       string `json:"receivedateformat"`
		Receivedate             string `json:"receivedate"`
		Receiptdateformat       string `json:"receiptdateformat"`
		Receiptdate             string `json:"receiptdate"`
		Fulfillexpeditecriteria string `json:"fulfillexpeditecriteria"`
		Companynumb             string `json:"companynumb"`
		Duplicate               string `json:"duplicate"`
		Reportduplicate         struct {
			Duplicatesource string `json:"duplicatesource"`
			Duplicatenumb   string `json:"duplicatenumb"`
		} `json:"reportduplicate"`
		Primarysource struct {
			Reportercountry string `json:"reportercountry"`
			Qualification   string `json:"qualification"`
		} `json:"primarysource"`
		Sender struct {
			Sendertype         string `json:"sendertype"`
			Senderorganization string `json:"senderorganization"`
		} `json:"sender"`
		Receiver struct {
			Receivertype         string `json:"receivertype"`
			Receiverorganization string `json:"receiverorganization"`
		} `json:"receiver"`
		Patient struct {
			Patientonsetage     string `json:"patientonsetage"`
			Patientonsetageunit string `json:"patientonsetageunit"`
			Patientweight       string `json:"patientweight"`
			Patientsex          string `json:"patientsex"`
			Reaction            []struct {
				Reactionmeddraversionpt string `json:"reactionmeddraversionpt"`
				Reactionmeddrapt        string `json:"reactionmeddrapt"`
				Reactionoutcome         string `json:"reactionoutcome"`
			} `json:"reaction"`
			Drug []struct {
				Drugcharacterization         string `json:"drugcharacterization"`
				Medicinalproduct             string `json:"medicinalproduct"`
				Drugbatchnumb                string `json:"drugbatchnumb,omitempty"`
				Drugstructuredosagenumb      string `json:"drugstructuredosagenumb,omitempty"`
				Drugstructuredosageunit      string `json:"drugstructuredosageunit,omitempty"`
				Drugseparatedosagenumb       string `json:"drugseparatedosagenumb,omitempty"`
				Drugintervaldosageunitnumb   string `json:"drugintervaldosageunitnumb,omitempty"`
				Drugintervaldosagedefinition string `json:"drugintervaldosagedefinition,omitempty"`
				Drugadministrationroute      string `json:"drugadministrationroute"`
				Drugindication               string `json:"drugindication"`
				Actiondrug                   string `json:"actiondrug,omitempty"`
				Drugrecurreadministration    string `json:"drugrecurreadministration,omitempty"`
				Openfda                      struct {
					ApplicationNumber []string `json:"application_number"`
					BrandName         []string `json:"brand_name"`
					GenericName       []string `json:"generic_name"`
					ManufacturerName  []string `json:"manufacturer_name"`
					ProductNdc        []string `json:"product_ndc"`
					ProductType       []string `json:"product_type"`
					Route             []string `json:"route"`
					SubstanceName     []string `json:"substance_name"`
					Rxcui             []string `json:"rxcui"`
					SplID             []string `json:"spl_id"`
					SplSetID          []string `json:"spl_set_id"`
					PackageNdc        []string `json:"package_ndc"`
					Unii              []string `json:"unii"`
				} `json:"openfda,omitempty"`
				Openfda_2 struct {
					ApplicationNumber []string `json:"application_number"`
					BrandName         []string `json:"brand_name"`
					GenericName       []string `json:"generic_name"`
					ManufacturerName  []string `json:"manufacturer_name"`
					ProductNdc        []string `json:"product_ndc"`
					ProductType       []string `json:"product_type"`
					Route             []string `json:"route"`
					SubstanceName     []string `json:"substance_name"`
					Rxcui             []string `json:"rxcui"`
					SplID             []string `json:"spl_id"`
					SplSetID          []string `json:"spl_set_id"`
					PackageNdc        []string `json:"package_ndc"`
					Nui               []string `json:"nui"`
					PharmClassMoa     []string `json:"pharm_class_moa"`
					PharmClassEpc     []string `json:"pharm_class_epc"`
					Unii              []string `json:"unii"`
				} `json:"openfda,omitempty"`
				Openfda_3 struct {
					ApplicationNumber []string `json:"application_number"`
					BrandName         []string `json:"brand_name"`
					GenericName       []string `json:"generic_name"`
					ManufacturerName  []string `json:"manufacturer_name"`
					ProductNdc        []string `json:"product_ndc"`
					ProductType       []string `json:"product_type"`
					Route             []string `json:"route"`
					SubstanceName     []string `json:"substance_name"`
					Rxcui             []string `json:"rxcui"`
					SplID             []string `json:"spl_id"`
					SplSetID          []string `json:"spl_set_id"`
					PackageNdc        []string `json:"package_ndc"`
					Nui               []string `json:"nui"`
					PharmClassMoa     []string `json:"pharm_class_moa"`
					PharmClassPe      []string `json:"pharm_class_pe"`
					PharmClassCs      []string `json:"pharm_class_cs"`
					PharmClassEpc     []string `json:"pharm_class_epc"`
					Unii              []string `json:"unii"`
				} `json:"openfda,omitempty"`
				Openfda_4 struct {
				} `json:"openfda,omitempty"`
			} `json:"drug"`
		} `json:"patient"`
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

func query_construct() string {
	// query construct
	fmt.Println("Enter openFDA query: ")
	var query string
	fmt.Scanln(&query)
	var full_query string
	full_query = baseURL + query + "&limit=" + limit
	return full_query
}

func find_meta_data() (string, int, int) {
	meta_query := query_construct()
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
	meta_query, skips_required, limit_int := find_meta_data()
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

func main() {

	data := get_data()

	csvFile, err := os.Create("./output_data/openFDA_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	// Define header row
	headerRow := []string{
		"date\tpatient_sex\tage\tweight\tgeneric_name\t",
	}
	writer.Write(headerRow)

	for _, data_page := range data {
		for _, usance := range data_page.Results {
			writer.Comma = '\t'
			var row []string
			row = append(row, usance.Receiptdate)
			row = append(row, usance.Patient.Patientsex)
			row = append(row, usance.Patient.Patientonsetage)
			row = append(row, usance.Patient.Patientweight)
			row = append(row, strings.Join(usance.Patient.Drug[0].Openfda.GenericName, "|"))

			writer.Write(row)
			writer.Flush() // Data flush
		}
	}
}
