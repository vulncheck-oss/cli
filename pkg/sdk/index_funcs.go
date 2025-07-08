package sdk

import (
	"encoding/json"

	"net/http"
	"net/url"

	"github.com/vulncheck-oss/cli/pkg/client"
)

// THIS FILE IS GENERATED. DO NOT EDIT

type Index7zipResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySevenZip `json:"data"`
}

// GetIndex7zip -  7Zip security advisories are official notifications released by 7Zip to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndex7zip(queryParameters ...IndexQueryParameters) (responseJSON *Index7zipResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("7zip"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexA10Response struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryA10 `json:"data"`
}

// GetIndexA10 -  A10 Networks security advisories are official notifications released by A10 Networks to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexA10(queryParameters ...IndexQueryParameters) (responseJSON *IndexA10Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("a10"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAbbResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryABBAdvisory `json:"data"`
}

// GetIndexAbb -  ABB vulnerabilities refer to security flaws that can be exploited in products and systems developed by ABB, a multinational technology company. These vulnerabilities can potentially lead to unauthorized access, manipulation of data, and disruption of critical infrastructure.

func (c *Client) GetIndexAbb(queryParameters ...IndexQueryParameters) (responseJSON *IndexAbbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("abb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAbbottResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryAbbott `json:"data"`
}

// GetIndexAbbott -  Abbott product advisories are official notifications released by Abbott to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAbbott(queryParameters ...IndexQueryParameters) (responseJSON *IndexAbbottResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("abbott"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAbsoluteResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAbsolute `json:"data"`
}

// GetIndexAbsolute -  Absolute security advisories are official notifications released by Absolute to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAbsolute(queryParameters ...IndexQueryParameters) (responseJSON *IndexAbsoluteResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("absolute"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAcronisResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAcronis `json:"data"`
}

// GetIndexAcronis -  Acronis security advisories are official notifications released by Acronis to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAcronis(queryParameters ...IndexQueryParameters) (responseJSON *IndexAcronisResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("acronis"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAdobeResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAdobeAdvisory `json:"data"`
}

// GetIndexAdobe -  Adobe Security Bulletins are official notifications released by Adobe Systems to address security vulnerabilities and updates in their software products. These bulletins provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAdobe(queryParameters ...IndexQueryParameters) (responseJSON *IndexAdobeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("adobe"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAixResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAIX `json:"data"`
}

// GetIndexAix -  AIX security advisories are official notifications released by IBM to address security vulnerabilities and updates in the AIX operating system. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexAix(queryParameters ...IndexQueryParameters) (responseJSON *IndexAixResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aix"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAlephResearchResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAlephResearch `json:"data"`
}

// GetIndexAlephResearch -  Aleph Research Vulnerability Reports are official notifications released by Aleph Research, a part of HCL Technologies, to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAlephResearch(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlephResearchResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aleph-research"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAlmaResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryAlmaLinuxUpdate `json:"data"`
}

// GetIndexAlma -  AlmaLinux is a popular community-driven Linux distribution that is built as a replacement for CentOS, which was recently discontinued by Red Hat. Like any other operating system, AlmaLinux is not immune to vulnerabilities and security flaws. Errata vulnerabilities refer to security issues that have been identified in a software system and require a patch or update to fix them. AlmaLinux has a dedicated team that constantly monitors for errata vulnerabilities and releases patches and updates to ensure that the system remains secure.

func (c *Client) GetIndexAlma(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlmaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("alma"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAlpineResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryAlpineLinuxSecDB `json:"data"`
}

// GetIndexAlpine -  The Alpine Linux Security Database is a public repository that maintains a comprehensive list of security vulnerabilities that have been identified in the Alpine Linux distribution. This database is an essential resource for Alpine Linux users who want to stay informed about potential security threats and vulnerabilities. The database provides detailed information about each security issue, including its severity level, affected components, and recommended fixes. Additionally, the Alpine Linux Security Team regularly updates the database with new vulnerabilities and patches, ensuring that users have access to the latest information and recommendations for securing their systems. The Alpine Linux Security Database is a critical component of the distribution's security infrastructure, and its transparency and accessibility reflect the project's commitment to ensuring the safety and reliability of its users' systems.

func (c *Client) GetIndexAlpine(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlpineResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("alpine"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAlpinePurlsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.PurlsPurlResponse `json:"data"`
}

// GetIndexAlpinePurls -  Alpine purls is a collection of Alpine package purls with their associated versions and cves.

func (c *Client) GetIndexAlpinePurls(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlpinePurlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("alpine-purls"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAmazonResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUpdate `json:"data"`
}

// GetIndexAmazon -  The Amazon Linux Security Center is a dedicated portal that provides users of Amazon Linux with a central location for information related to security on the platform. The security center includes access to documentation, guidance, and best practices to help users configure and secure their Amazon Linux environments. The center also provides access to the Amazon Linux AMI vulnerability database, which lists all known security vulnerabilities affecting the operating system, as well as information on how to mitigate each vulnerability.

func (c *Client) GetIndexAmazon(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmazonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("amazon"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAmazonCveResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryAmazonCVE `json:"data"`
}

// GetIndexAmazonCve -  Amazon CVEs are official notifications released by AWS to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexAmazonCve(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmazonCveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("amazon-cve"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAmdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAMD `json:"data"`
}

// GetIndexAmd -  AMD security bulletins are official notifications released by AMD to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAmd(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("amd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAmiResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAMI `json:"data"`
}

// GetIndexAmi -  AMI security advisories are official notifications released by the AMI Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAmi(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ami"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAnchoreNvdOverrideResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryAnchoreNVDOverride `json:"data"`
}

// GetIndexAnchoreNvdOverride -  Anchore NVD Data Overrides is an index of data overrides for the NVD dataset curated by Anchore that provides additional data that might be missing from NVD.

func (c *Client) GetIndexAnchoreNvdOverride(queryParameters ...IndexQueryParameters) (responseJSON *IndexAnchoreNvdOverrideResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("anchore-nvd-override"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAndroidResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryAndroidAdvisory `json:"data"`
}

// GetIndexAndroid -  Android security bulletins are official notifications released by Google to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAndroid(queryParameters ...IndexQueryParameters) (responseJSON *IndexAndroidResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("android"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheActivemqResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryApacheActiveMQ `json:"data"`
}

// GetIndexApacheActivemq -  Apache ActiveMQ security advisories are official notifications released by the open source Apache ActiveMQ project to address security vulnerabilities and updates in the open source Apache ActiveMQ project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheActivemq(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheActivemqResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-activemq"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheArchivaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheArchiva `json:"data"`
}

// GetIndexApacheArchiva -  Apache Archiva security vulnerabilities are official notifications released by the open source Apache Archiva project to address security vulnerabilities and updates in the open source Apache Archiva project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheArchiva(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheArchivaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-archiva"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheArrowResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheArrow `json:"data"`
}

// GetIndexApacheArrow -  Apache Arrow security issues are official notifications released by the open source Apache Arrow project to address security vulnerabilities and updates in the open source Apache Arrow project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheArrow(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheArrowResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-arrow"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheCamelResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheCamel `json:"data"`
}

// GetIndexApacheCamel -  Apache Camel security advisories are official notifications released by the open source Apache Camel project to address security vulnerabilities and updates in the open source Apache Camel project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheCamel(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheCamelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-camel"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheCommonsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheCommons `json:"data"`
}

// GetIndexApacheCommons -  Apache Commons security vulnerabilities are official notifications released by the open source Apache Commons project to address security vulnerabilities and updates in the open source Apache Commons project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheCommons(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheCommonsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-commons"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheCouchdbResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheCouchDB `json:"data"`
}

// GetIndexApacheCouchdb -  Apache CouchDB security issues are official notifications released by the open source Apache CouchDB project to address security vulnerabilities and updates in the open source Apache CouchDB project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheCouchdb(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheCouchdbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-couchdb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheFlinkResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheFlink `json:"data"`
}

// GetIndexApacheFlink -  Apache Flink security updates are official notifications released by the open source Apache Flink project to address security vulnerabilities and updates in the open source Apache Flink project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheFlink(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheFlinkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-flink"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheGuacamoleResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryApacheGuacamole `json:"data"`
}

// GetIndexApacheGuacamole -  Apache Guacamole security reports are official notifications released by the open source Apache Guacamole project to address security vulnerabilities and updates in the open source Apache Guacamole project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheGuacamole(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheGuacamoleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-guacamole"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheHadoopResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheHadoop `json:"data"`
}

// GetIndexApacheHadoop -  Apache Hadoop CVEs are official notifications released by the open source Apache Hadoop project to address security vulnerabilities and updates in the open source Apache Hadoop project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheHadoop(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheHadoopResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-hadoop"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheHttpResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryApacheHTTP `json:"data"`
}

// GetIndexApacheHttp -  Apache HTTP security vulnerabilities are official notifications released by the open source Apache project to address security vulnerabilities and updates in the open source Apache HTTP project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheHttp(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheHttpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-http"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheJspwikiResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheJSPWiki `json:"data"`
}

// GetIndexApacheJspwiki -  Apache JSPWiki CVEs are official notifications released by the open source Apache JSPWiki project to address security vulnerabilities and updates in the open source Apache OpenMeetings project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheJspwiki(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheJspwikiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-jspwiki"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheKafkaResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheKafka `json:"data"`
}

// GetIndexApacheKafka -  Apache Kafka security vulnerabilities are official notifications released by the open source Apache Kafka project to address security vulnerabilities and updates in the open source Apache Kafka project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheKafka(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheKafkaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-kafka"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheLoggingservicesResponse struct {
	Benchmark float64                                `json:"_benchmark"`
	Meta      IndexMeta                              `json:"_meta"`
	Data      []client.AdvisoryApacheLoggingServices `json:"data"`
}

// GetIndexApacheLoggingservices -  Apache Logging Services known vulnerabilities are official notifications released by the open source Apache Logging Services project to address security vulnerabilities and updates in the open source Apache Logging Services project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheLoggingservices(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheLoggingservicesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-loggingservices"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheNifiResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryApacheNiFi `json:"data"`
}

// GetIndexApacheNifi -  Apache NiFi security vulnerabilities are official notifications released by the open source Apache NiFi project to address security vulnerabilities and updates in the open source Apache NiFi project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheNifi(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheNifiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-nifi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheOfbizResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheOFBiz `json:"data"`
}

// GetIndexApacheOfbiz -  Apache OFBiz security vulnerabilities are official notifications released by the open source Apache OFBiz project to address security vulnerabilities and updates in the open source Apache OFBiz project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheOfbiz(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheOfbizResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-ofbiz"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheOpenmeetingsResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryApacheOpenMeetings `json:"data"`
}

// GetIndexApacheOpenmeetings -  Apache OpenMeetings security vulnerabilities are official notifications released by the open source Apache OpenMeetings project to address security vulnerabilities and updates in the open source Apache OpenMeetings project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheOpenmeetings(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheOpenmeetingsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-openmeetings"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheOpenofficeResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryApacheOpenOffice `json:"data"`
}

// GetIndexApacheOpenoffice -  Apache OpenOffice security bulletins are official notifications released by the open source Apache OpenOffice project to address security vulnerabilities and updates in the open source Apache OpenOffice project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheOpenoffice(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheOpenofficeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-openoffice"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApachePulsarResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApachePulsar `json:"data"`
}

// GetIndexApachePulsar -  Apache Pulsar security advisories are official notifications released by the open source Apache Pulsar project to address security vulnerabilities and updates in the open source Apache Pulsar project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApachePulsar(queryParameters ...IndexQueryParameters) (responseJSON *IndexApachePulsarResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-pulsar"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheShiroResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheShiro `json:"data"`
}

// GetIndexApacheShiro -  Apache Shiro vulnerability reports are official notifications released by the open source Apache Shiro project to address security vulnerabilities and updates in the open source Apache Shiro project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheShiro(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheShiroResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-shiro"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheSparkResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheSpark `json:"data"`
}

// GetIndexApacheSpark -  Apache Spark cves are official notifications released by the open source Apache Spark project to address security vulnerabilities and updates in the open source Apache Spark project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheSpark(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheSparkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-spark"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheStrutsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheStruts `json:"data"`
}

// GetIndexApacheStruts -  Apache Struts security bulletins are official notifications released by the open source Apache Struts project to address security vulnerabilities and updates in the open source Apache Struts project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheStruts(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheStrutsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-struts"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheSubversionResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryApacheSubversion `json:"data"`
}

// GetIndexApacheSubversion -  Apache Subversion security advisories are official notifications released by the open source Apache Subversion project to address security vulnerabilities and updates in the open source Apache Subversion project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheSubversion(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheSubversionResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-subversion"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheSupersetResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryApacheSuperset `json:"data"`
}

// GetIndexApacheSuperset -  Apache Superset cves are official notifications released by the open source Apache Superset project to address security vulnerabilities and updates in the open source Apache Superset project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheSuperset(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheSupersetResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-superset"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheTomcatResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheTomcat `json:"data"`
}

// GetIndexApacheTomcat -  Apache Tomcat security vunlnerabilities are official notifications released by the open source Apache Struts project to address security vulnerabilities and updates in the open source Apache Strus project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheTomcat(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheTomcatResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-tomcat"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexApacheZookeeperResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryApacheZooKeeper `json:"data"`
}

// GetIndexApacheZookeeper -  Apache ZooKeeper vulnerability reports are official notifications released by the open source Apache ZooKeeper project to address vulnerabilities and updates in the open source Apache ZooKeeper project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexApacheZookeeper(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheZookeeperResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-zookeeper"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAppcheckResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAppCheck `json:"data"`
}

// GetIndexAppcheck -  AppCheck security alerts are official notifications released by AppCheck to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAppcheck(queryParameters ...IndexQueryParameters) (responseJSON *IndexAppcheckResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("appcheck"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAppgateResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAppgate `json:"data"`
}

// GetIndexAppgate -  Appgate SDP security advisories sare official notifications released by Appgate to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAppgate(queryParameters ...IndexQueryParameters) (responseJSON *IndexAppgateResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("appgate"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAppleResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAppleAdvisory `json:"data"`
}

// GetIndexApple -  Apple regularly releases security updates to address vulnerabilities in its operating systems, software applications, and devices. These updates are critical for maintaining the security of Apple products and protecting users from potential cyber threats. Apple encourages users to promptly install security updates to ensure that their devices are protected against known vulnerabilities and to stay vigilant against potential new threats. category: Product Security Advisories

func (c *Client) GetIndexApple(queryParameters ...IndexQueryParameters) (responseJSON *IndexAppleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apple"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexArchResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryArchIssue `json:"data"`
}

// GetIndexArch -  Arch Linux's rolling-release model ensures that security patches are promptly released and distributed to users, minimizing the exposure to known vulnerabilities and providing a relatively secure system when kept up to date.

func (c *Client) GetIndexArch(queryParameters ...IndexQueryParameters) (responseJSON *IndexArchResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("arch"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAristaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryArista `json:"data"`
}

// GetIndexArista -  Arista Networks security advisories are official notifications released by the Arista Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexArista(queryParameters ...IndexQueryParameters) (responseJSON *IndexAristaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("arista"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexArubaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAruba `json:"data"`
}

// GetIndexAruba -  Aruba security advisories are official notifications released by Aruba’s Security Incident Response Team (SIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAruba(queryParameters ...IndexQueryParameters) (responseJSON *IndexArubaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aruba"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAsrgResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryASRG `json:"data"`
}

// GetIndexAsrg -  Automotive Security Research Group (ASRG) security advisories are official notifications released by ASRG to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAsrg(queryParameters ...IndexQueryParameters) (responseJSON *IndexAsrgResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("asrg"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAssetnoteResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryAssetNote `json:"data"`
}

// GetIndexAssetnote -  AssetNote security advisories are official notifications released by AssetNote to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAssetnote(queryParameters ...IndexQueryParameters) (responseJSON *IndexAssetnoteResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("assetnote"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAsteriskResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAsterisk `json:"data"`
}

// GetIndexAsterisk -  Asterisk security advisories are official notifications released by the open source Asterisk project to address security vulnerabilities and updates in the open source Asterisk project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAsterisk(queryParameters ...IndexQueryParameters) (responseJSON *IndexAsteriskResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("asterisk"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAstraResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAstra `json:"data"`
}

// GetIndexAstra -  Astra security bulletins are official notifications released by Astra to address security vulnerabilities and updates for the Astra linux distrubution. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexAstra(queryParameters ...IndexQueryParameters) (responseJSON *IndexAstraResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("astra"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAsusResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAsus `json:"data"`
}

// GetIndexAsus -  Asus security advisories are official notifications released by Asus to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAsus(queryParameters ...IndexQueryParameters) (responseJSON *IndexAsusResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("asus"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAtlassianResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryAtlassianAdvisory `json:"data"`
}

// GetIndexAtlassian -  Atlassian security advisories are official notifications released by Atlassian to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems. Security advisories for Atlassian server products are released every Wednesday.

func (c *Client) GetIndexAtlassian(queryParameters ...IndexQueryParameters) (responseJSON *IndexAtlassianResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("atlassian"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAtlassianVulnsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAtlassianVuln `json:"data"`
}

// GetIndexAtlassianVulns -  Atlassian vulnerabilities are official notifications released by Atlassian to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems. Security advisories for Atlassian server products are released every Wednesday.

func (c *Client) GetIndexAtlassianVulns(queryParameters ...IndexQueryParameters) (responseJSON *IndexAtlassianVulnsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("atlassian-vulns"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAtredisResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAtredis `json:"data"`
}

// GetIndexAtredis -  Atredis Partners security advisories are official notifications released by Atredis Partners to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAtredis(queryParameters ...IndexQueryParameters) (responseJSON *IndexAtredisResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("atredis"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAuscertResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAusCert `json:"data"`
}

// GetIndexAuscert -  AusCERT Bulletins are periodic publications issued by AusCERT to inform their members about the latest cybersecurity threats, vulnerabilities, and incidents. These bulletins provide concise summaries, technical details, and recommended actions to mitigate risks and protect systems and networks. They serve as valuable resources for organizations seeking up-to-date information and guidance to enhance their security defenses.

func (c *Client) GetIndexAuscert(queryParameters ...IndexQueryParameters) (responseJSON *IndexAuscertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("auscert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAutodeskResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAutodesk `json:"data"`
}

// GetIndexAutodesk -  Autodesk security advisories are official notifications released by Autodesk to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAutodesk(queryParameters ...IndexQueryParameters) (responseJSON *IndexAutodeskResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("autodesk"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAvayaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAvaya `json:"data"`
}

// GetIndexAvaya -  Avaya security advisories are official notifications released by Avaya to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAvaya(queryParameters ...IndexQueryParameters) (responseJSON *IndexAvayaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("avaya"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAvevaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAVEVAAdvisory `json:"data"`
}

// GetIndexAveva -  Aveva security advisories are official notifications released by Aveva to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAveva(queryParameters ...IndexQueryParameters) (responseJSON *IndexAvevaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aveva"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAvigilonResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAvigilon `json:"data"`
}

// GetIndexAvigilon -  Avigilon security advisories are official notifications released by Avigilon to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAvigilon(queryParameters ...IndexQueryParameters) (responseJSON *IndexAvigilonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("avigilon"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAwsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAWS `json:"data"`
}

// GetIndexAws -  AWS security bulletins are official notifications released by Amazon Web Services to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAws(queryParameters ...IndexQueryParameters) (responseJSON *IndexAwsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aws"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAxisResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAxis `json:"data"`
}

// GetIndexAxis -  Axis OS security advisories are official notifications released by Axis to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAxis(queryParameters ...IndexQueryParameters) (responseJSON *IndexAxisResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("axis"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAzulResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAzul `json:"data"`
}

// GetIndexAzul -  Azul Common Vulnerabilities and Exposures are official notifications released by Azul to address security vulnerabilities and updates in their software and hardware products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexAzul(queryParameters ...IndexQueryParameters) (responseJSON *IndexAzulResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("azul"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBandrResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryBandr `json:"data"`
}

// GetIndexBandr -  B&R Security Bulletins are regular notifications released by B&R Industrial Automation, a leading provider of automation solutions. These bulletins aim to address security vulnerabilities and provide updates related to B&R's products and software. They offer important information on potential risks, recommended patches or updates, and best practices to enhance the security of B&R automation systems deployed in various industries.

func (c *Client) GetIndexBandr(queryParameters ...IndexQueryParameters) (responseJSON *IndexBandrResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bandr"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBaxterResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBaxterAdvisory `json:"data"`
}

// GetIndexBaxter -  Baxter Security Advisories are official notifications issued by Baxter International, a global healthcare company, to address security vulnerabilities and updates in their medical devices and software. These advisories inform healthcare professionals and users about potential risks, recommended actions, and available patches or updates to ensure the security and integrity of Baxter's products. They play a crucial role in promoting patient safety and guiding healthcare organizations in implementing necessary security measures.

func (c *Client) GetIndexBaxter(queryParameters ...IndexQueryParameters) (responseJSON *IndexBaxterResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("baxter"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBbraunResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBBraunAdvisory `json:"data"`
}

// GetIndexBbraun -  BBraun security advisories are official notifications released by BBraun to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexBbraun(queryParameters ...IndexQueryParameters) (responseJSON *IndexBbraunResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bbraun"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBdResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryBectonDickinsonAdvisory `json:"data"`
}

// GetIndexBd -  The `bd` index contains data on advisories published by Becton Dickinson. Becton Dickinson is a medical technology company that develops, manufactures, and sells medical devices, instrument systems, and reagents. The company is headquartered in Franklin Lakes, New Jersey, United States.

func (c *Client) GetIndexBd(queryParameters ...IndexQueryParameters) (responseJSON *IndexBdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBduResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBDUAdvisory `json:"data"`
}

// GetIndexBdu -  The `bdu` index contains security advisories that are official communications issued by military or government agencies to provide information, guidance, and updates related to security risks and threats. These advisories are designed to provide personnel with essential information and recommendations to minimize the risk of security incidents and protect critical assets.

func (c *Client) GetIndexBdu(queryParameters ...IndexQueryParameters) (responseJSON *IndexBduResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bdu"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBeckhoffResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryBeckhoffAdvisory `json:"data"`
}

// GetIndexBeckhoff -  Beckhoff Advisories are security notifications issued by Beckhoff Automation, a prominent provider of automation technology. These advisories inform customers and users about potential vulnerabilities, patches, and mitigations related to Beckhoff's hardware, software, and industrial control systems. They provide essential information and guidance to help organizations protect their automation infrastructure and ensure the secure operation of their Beckhoff-based systems.

func (c *Client) GetIndexBeckhoff(queryParameters ...IndexQueryParameters) (responseJSON *IndexBeckhoffResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("beckhoff"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBeldenResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBeldenAdvisory `json:"data"`
}

// GetIndexBelden -  Belden Security Bulletins are regular notifications issued by Belden Inc., a global leader in signal transmission solutions. These bulletins provide updates, advisories, and recommendations related to the security of Belden's products and systems, including network infrastructure, industrial control systems, and data centers. They serve as a valuable resource for Belden customers and users to stay informed about potential vulnerabilities, best practices, and available patches or updates to ensure the security and reliability of their communication networks.

func (c *Client) GetIndexBelden(queryParameters ...IndexQueryParameters) (responseJSON *IndexBeldenResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("belden"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBeyondTrustResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBeyondTrust `json:"data"`
}

// GetIndexBeyondTrust -  Beyond Trust security advisories are official notifications released by Beyond Trust to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexBeyondTrust(queryParameters ...IndexQueryParameters) (responseJSON *IndexBeyondTrustResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("beyond-trust"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBinarlyResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryBinarly `json:"data"`
}

// GetIndexBinarly -  Binarly advisories are official notifications released by Binarly to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexBinarly(queryParameters ...IndexQueryParameters) (responseJSON *IndexBinarlyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("binarly"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBitdefenderResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBitDefender `json:"data"`
}

// GetIndexBitdefender -  Bitdefender security advisories are official notifications released by Bitdefender to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexBitdefender(queryParameters ...IndexQueryParameters) (responseJSON *IndexBitdefenderResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bitdefender"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBlackberryResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryBlackBerry `json:"data"`
}

// GetIndexBlackberry -  BlackBerry security advisories are official notifications released by the BlackBerry Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexBlackberry(queryParameters ...IndexQueryParameters) (responseJSON *IndexBlackberryResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("blackberry"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBlsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryBLS `json:"data"`
}

// GetIndexBls -  Black Lantern security advisories are official notifications released by Black Lantern to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexBls(queryParameters ...IndexQueryParameters) (responseJSON *IndexBlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bls"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBoschResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryBoschAdvisory `json:"data"`
}

// GetIndexBosch -  Bosch Security Advisories are official notifications released by Bosch, a renowned technology company, to address security vulnerabilities and updates in their security products and solutions. These advisories provide detailed information on identified vulnerabilities, potential risks, and recommended actions to mitigate security threats. By promptly informing customers and users about vulnerabilities and offering guidance, Bosch Security Advisories help maintain the integrity and resilience of their security systems and protect against potential cyberattacks.

func (c *Client) GetIndexBosch(queryParameters ...IndexQueryParameters) (responseJSON *IndexBoschResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bosch"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBostonScientificResponse struct {
	Benchmark float64                                   `json:"_benchmark"`
	Meta      IndexMeta                                 `json:"_meta"`
	Data      []client.AdvisoryBostonScientificAdvisory `json:"data"`
}

// GetIndexBostonScientific -  Boston Scientific Advisories are official notifications released by Boston Scientific Corporation, a global medical technology company. These advisories inform healthcare professionals and users about important updates, safety concerns, and recommended actions related to Boston Scientific medical devices and therapies. They play a critical role in ensuring patient safety and guiding healthcare providers in implementing necessary measures to address potential risks and maintain the proper functioning of Boston Scientific products.

func (c *Client) GetIndexBostonScientific(queryParameters ...IndexQueryParameters) (responseJSON *IndexBostonScientificResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("boston-scientific"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexBotnetsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryBotnet `json:"data"`
}

// GetIndexBotnets -  The VulnCheck Botnets index contains data related to various botnets. The index contains listings of botnets and citations for the CVE they have been known to use.

func (c *Client) GetIndexBotnets(queryParameters ...IndexQueryParameters) (responseJSON *IndexBotnetsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("botnets"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCaCyberCentreResponse struct {
	Benchmark float64                                `json:"_benchmark"`
	Meta      IndexMeta                              `json:"_meta"`
	Data      []client.AdvisoryCACyberCentreAdvisory `json:"data"`
}

// GetIndexCaCyberCentre -  The Cyber Centre issues alerts and advisories on potential, imminent or actual cyber threats, vulnerabilities or incidents affecting Canada's critical infrastructure.

func (c *Client) GetIndexCaCyberCentre(queryParameters ...IndexQueryParameters) (responseJSON *IndexCaCyberCentreResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ca-cyber-centre"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCanvasResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryCanvasExploit `json:"data"`
}

// GetIndexCanvas -  CANVAS Exploit Packs developed by Gleg are powerful tools used in penetration testing and vulnerability assessment. These exploit packs provide a comprehensive range of exploits and attack vectors to assess the security of computer systems and applications.

func (c *Client) GetIndexCanvas(queryParameters ...IndexQueryParameters) (responseJSON *IndexCanvasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("canvas"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCarestreamResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryCarestreamAdvisory `json:"data"`
}

// GetIndexCarestream -  Carestream Product Security Advisories are official notifications released by Carestream Health, a leading provider of medical imaging and healthcare IT solutions. These advisories address security vulnerabilities and updates related to Carestream's products and software in the healthcare industry. They provide essential information, including the nature of the vulnerability, potential risks, recommended actions, and available patches or updates to mitigate security risks and ensure the confidentiality, integrity, and availability of patient data and healthcare systems. Carestream Product Security Advisories are crucial in helping healthcare organizations maintain a secure and protected environment for patient care.

func (c *Client) GetIndexCarestream(queryParameters ...IndexQueryParameters) (responseJSON *IndexCarestreamResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("carestream"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCargoResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexCargo -  Cargo (Rust) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexCargo(queryParameters ...IndexQueryParameters) (responseJSON *IndexCargoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cargo"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCarrierResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryCarrier `json:"data"`
}

// GetIndexCarrier -  Carrier product security advisories are official notifications released by the Carrier Product Security Incident Response Team (Carrier PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCarrier(queryParameters ...IndexQueryParameters) (responseJSON *IndexCarrierResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("carrier"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCblMarinerResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryCBLMariner `json:"data"`
}

// GetIndexCblMariner -  CBL-Mariner contains vulnerabilities detected in the Microsoft CBL Mariner linux distribution.

func (c *Client) GetIndexCblMariner(queryParameters ...IndexQueryParameters) (responseJSON *IndexCblMarinerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cbl-mariner"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCentosResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCESA `json:"data"`
}

// GetIndexCentos -  CentOS Security Advisories are official notifications issued by the CentOS project, a popular open-source Linux distribution. These advisories provide information on security vulnerabilities, patches, and updates relevant to CentOS operating systems. They help CentOS users stay informed about potential risks, recommended actions, and available fixes to maintain the security and stability of their CentOS-based systems. CentOS Security Advisories play a vital role in assisting system administrators and users in effectively managing and securing their CentOS deployments.

func (c *Client) GetIndexCentos(queryParameters ...IndexQueryParameters) (responseJSON *IndexCentosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("centos"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCertBeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertBE `json:"data"`
}

// GetIndexCertBe -  CERT BE security advisories are official notifications released by the Centre for CyberSecurity Belgium to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexCertBe(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertBeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-be"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCertInResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertIN `json:"data"`
}

// GetIndexCertIn -  CERT IN security advisories are official notifications released by India's national CERT (Computer Emergency Response Team) to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexCertIn(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertInResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-in"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCertIrSecurityAlertsResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryCertIRSecurityAlert `json:"data"`
}

// GetIndexCertIrSecurityAlerts -  CERT IR security warnings are official notifications released by Iran's national CERT (Computer Emergency Response Team) to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexCertIrSecurityAlerts(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertIrSecurityAlertsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-ir-security-alerts"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCertSeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertSE `json:"data"`
}

// GetIndexCertSe -  CERT SE security advisories are official notifications released by Sweden's national CSIRT (Computer Security Incident Response Team) to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexCertSe(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertSeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-se"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCertUaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertUA `json:"data"`
}

// GetIndexCertUa -  CERT UA security advisories are official notifications released by the Ukraine CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexCertUa(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertUaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-ua"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCerteuResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCERTEUAdvisory `json:"data"`
}

// GetIndexCerteu -  Cert-EU Bulletins are periodic publications issued by Cert-EU to inform their members about the latest cybersecurity threats, vulnerabilities, and incidents. These bulletins provide concise summaries, technical details, and recommended actions to mitigate risks and protect systems and networks.

func (c *Client) GetIndexCerteu(queryParameters ...IndexQueryParameters) (responseJSON *IndexCerteuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("certeu"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCertfrResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCertFRAdvisory `json:"data"`
}

// GetIndexCertfr -  CERT-FR security alerts are official notifications released by the French national and governmental CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexCertfr(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertfrResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("certfr"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexChainguardResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryChainGuard `json:"data"`
}

// GetIndexChainguard -  ChainGuard is an enterprise Linux undistribution based on Wolfi that combines the best aspects of existing container base images with default security measures that will include software signatures powered by Sigstore, provenance, and software bills of material (SBOM).

func (c *Client) GetIndexChainguard(queryParameters ...IndexQueryParameters) (responseJSON *IndexChainguardResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("chainguard"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCheckpointResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryCheckPoint `json:"data"`
}

// GetIndexCheckpoint -  CheckPoint security advisories are official notifications released by CheckPoint to address security vulnerabilities and updates in the third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCheckpoint(queryParameters ...IndexQueryParameters) (responseJSON *IndexCheckpointResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("checkpoint"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexChromeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryChrome `json:"data"`
}

// GetIndexChrome -  Chrome release updates are periodic publications issued by the Google Chrome team to inform their members about the latest cybersecurity threats, vulnerabilities, and incidents. These bulletins provide concise summaries, technical details, and recommended actions to mitigate risks and protect systems and networks.

func (c *Client) GetIndexChrome(queryParameters ...IndexQueryParameters) (responseJSON *IndexChromeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("chrome"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCisaAlertsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCISAAlert `json:"data"`
}

// GetIndexCisaAlerts -  CISA (Cybersecurity and Infrastructure Security Agency) Alerts are official notifications issued by the United States' primary federal agency responsible for cybersecurity. These alerts provide timely and actionable information on emerging cyber threats, vulnerabilities, and incidents affecting critical infrastructure sectors. CISA Alerts offer guidance, recommended mitigation measures, and best practices to enhance the security and resilience of organizations, promoting a proactive approach to protecting critical systems and networks from cyber threats.

func (c *Client) GetIndexCisaAlerts(queryParameters ...IndexQueryParameters) (responseJSON *IndexCisaAlertsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisa-alerts"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCisaKevResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryKEVCatalogVulnerability `json:"data"`
}

// GetIndexCisaKev -  The CISA Known Exploited Vulnerabilities catalog contains a list of exploited vulnerabilities known to CISA.

func (c *Client) GetIndexCisaKev(queryParameters ...IndexQueryParameters) (responseJSON *IndexCisaKevResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisa-kev"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCiscoResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryCiscoAdvisory `json:"data"`
}

// GetIndexCisco -  Cisco security advisories are official notifications released by Cisco to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCisco(queryParameters ...IndexQueryParameters) (responseJSON *IndexCiscoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisco"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCiscoCsafResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCiscoCSAF `json:"data"`
}

// GetIndexCiscoCsaf -  Cisco CSAF is an index of Cisco security advisories in CSAF format.
func (c *Client) GetIndexCiscoCsaf(queryParameters ...IndexQueryParameters) (responseJSON *IndexCiscoCsafResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisco-csaf"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCiscoTalosResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryTalosAdvisory `json:"data"`
}

// GetIndexCiscoTalos -  The `cisco-talos` Security Advisories are official notifications released by the Talos research group within Cisco that provide information and updates on potential security vulnerabilities and threats affecting Cisco products and services.

func (c *Client) GetIndexCiscoTalos(queryParameters ...IndexQueryParameters) (responseJSON *IndexCiscoTalosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisco-talos"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCitrixResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCitrixAdvisory `json:"data"`
}

// GetIndexCitrix -  Citrix Security Advisories are official notifications released by Citrix Systems, a leading provider of digital workspace and networking solutions. These advisories address security vulnerabilities and updates in Citrix products, such as Citrix ADC, Citrix Gateway, and Citrix Virtual Apps and Desktops. They provide detailed information about the vulnerabilities, potential impact, and recommended actions, including patches or workarounds, to mitigate the risks. Citrix Security Advisories play a crucial role in helping organizations maintain the security and integrity of their Citrix deployments and protect against potential cyber threats.

func (c *Client) GetIndexCitrix(queryParameters ...IndexQueryParameters) (responseJSON *IndexCitrixResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("citrix"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexClarotyResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryClarotyVulnerability `json:"data"`
}

// GetIndexClaroty -  Team82 aligns with defenders of industrial, healthcare, and commercial networks, and provides indispensable threat and vulnerability research in order to ensure the safety, reliability, and integrity of systems within critical industries.

func (c *Client) GetIndexClaroty(queryParameters ...IndexQueryParameters) (responseJSON *IndexClarotyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("claroty"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCloudbeesResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCloudBees `json:"data"`
}

// GetIndexCloudbees -  CloudBees security advisories are official notifications released by CloudBees to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCloudbees(queryParameters ...IndexQueryParameters) (responseJSON *IndexCloudbeesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cloudbees"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCloudvulndbResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryCloudVulnDBAdvisory `json:"data"`
}

// GetIndexCloudvulndb -  CloudVulnDB is a comprehensive and continuously updated database that focuses on cataloging security vulnerabilities specific to cloud services and environments. It provides detailed information about vulnerabilities, including their impact, severity, affected platforms, and recommended mitigation strategies. CloudVulnDB serves as a valuable resource for security professionals and organizations seeking to proactively identify and address vulnerabilities in their cloud infrastructure, enabling them to enhance their overall security posture.

func (c *Client) GetIndexCloudvulndb(queryParameters ...IndexQueryParameters) (responseJSON *IndexCloudvulndbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cloudvulndb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCnnvdResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCNNVDEntryJSON `json:"data"`
}

// GetIndexCnnvd -  The Chinese National Vulnerability Database is one of two national vulnerability databases of the People’s Republic of China. It is operated by the China Information Technology Security Evaluation Center, the 13th Bureau of China’s foreign intelligence service, the Ministry of State Security.

func (c *Client) GetIndexCnnvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexCnnvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cnnvd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCnvdBulletinsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryCNVDBulletin `json:"data"`
}

// GetIndexCnvdBulletins -  The Chinese National Vulnerability Database (CNVD) is a service responsible for collecting and sharing information about software vulnerabilities that affect Chinese information systems. The CNVD publishes advisories about security flaws and vulnerabilities that have been identified in software products and systems.

func (c *Client) GetIndexCnvdBulletins(queryParameters ...IndexQueryParameters) (responseJSON *IndexCnvdBulletinsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cnvd-bulletins"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCnvdFlawsResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryCNVDFlaw `json:"data"`
}

// GetIndexCnvdFlaws -  The Chinese National Vulnerability Database (CNVD) is a service responsible for collecting and sharing information about software vulnerabilities that affect Chinese information systems. The CNVD publishes advisories about security flaws and vulnerabilities that have been identified in software products and systems.

func (c *Client) GetIndexCnvdFlaws(queryParameters ...IndexQueryParameters) (responseJSON *IndexCnvdFlawsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cnvd-flaws"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCocoapodsResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexCocoapods -  CocoaPods (Swift, Objective-C) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexCocoapods(queryParameters ...IndexQueryParameters) (responseJSON *IndexCocoapodsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cocoapods"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCodesysResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryCodesysAdvisory `json:"data"`
}

// GetIndexCodesys -  CODESYS Advisories are official notifications issued by CODESYS, a widely used development environment for programming industrial control systems. These advisories highlight security vulnerabilities, patches, and updates related to the CODESYS software. They provide important information on potential risks, recommended actions, and available fixes to address vulnerabilities and protect industrial automation systems from potential cyber threats. CODESYS Advisories help ensure the secure operation of control systems and assist system integrators and operators in maintaining the integrity and reliability of their industrial processes.

func (c *Client) GetIndexCodesys(queryParameters ...IndexQueryParameters) (responseJSON *IndexCodesysResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("codesys"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCommvaultResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCommVault `json:"data"`
}

// GetIndexCommvault -  CommVault security advisories are official notifications released by CommVault to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexCommvault(queryParameters ...IndexQueryParameters) (responseJSON *IndexCommvaultResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("commvault"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCompassSecurityResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryCompassSecurity `json:"data"`
}

// GetIndexCompassSecurity -  Compass Security advisories are official notifications released by Compass Security to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCompassSecurity(queryParameters ...IndexQueryParameters) (responseJSON *IndexCompassSecurityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("compass-security"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexComposerResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexComposer -  Composer (PHP) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexComposer(queryParameters ...IndexQueryParameters) (responseJSON *IndexComposerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("composer"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexConanResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexConan -  Conan (C/C++) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexConan(queryParameters ...IndexQueryParameters) (responseJSON *IndexConanResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("conan"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCoreimpactResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryCoreImpactExploit `json:"data"`
}

// GetIndexCoreimpact -  Core Impact is a library of expert validated exploits for safe and effective pen tests.
func (c *Client) GetIndexCoreimpact(queryParameters ...IndexQueryParameters) (responseJSON *IndexCoreimpactResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("coreimpact"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCrestronResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryCrestron `json:"data"`
}

// GetIndexCrestron -  Crestron security advisories are official notifications released by Crestron to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCrestron(queryParameters ...IndexQueryParameters) (responseJSON *IndexCrestronResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("crestron"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCurlResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCurl `json:"data"`
}

// GetIndexCurl -  Curl CVEs are official notifications released by the Curl open source project to address security vulnerabilities and updates in curl. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexCurl(queryParameters ...IndexQueryParameters) (responseJSON *IndexCurlResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("curl"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexCweResponse struct {
	Benchmark float64         `json:"_benchmark"`
	Meta      IndexMeta       `json:"_meta"`
	Data      []client.ApiCWE `json:"data"`
}

// GetIndexCwe -  The MITRE Common Weakness Enumeration (CWE) is a community-developed list of common software security weaknesses. The CWE is maintained by the MITRE Corporation, a not-for-profit organization that operates federally funded research and development centers (FFRDCs) sponsored by the U.S. government. The CWE is a valuable resource for software developers, security professionals, and other stakeholders in the software industry. It provides a standardized way to identify and describe common software security weaknesses, which helps to improve the security of software systems and applications.

func (c *Client) GetIndexCwe(queryParameters ...IndexQueryParameters) (responseJSON *IndexCweResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cwe"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDahuaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryDahua `json:"data"`
}

// GetIndexDahua -  Dahua security advisories are official notifications released by the Dahua Product Security Incident Response Team (Dahua PSIRT)  to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDahua(queryParameters ...IndexQueryParameters) (responseJSON *IndexDahuaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dahua"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDassaultResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryDassault `json:"data"`
}

// GetIndexDassault -  Dassault Systèmes security advisories are official notifications released by Dassault to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDassault(queryParameters ...IndexQueryParameters) (responseJSON *IndexDassaultResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dassault"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDebianResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryVulnerableDebianPackage `json:"data"`
}

// GetIndexDebian -  Debian Security Tracker - `debian-security-tracker` index is a service that provides information and updates on security vulnerabilities and issues affecting Debian packages and software. The Debian Security Tracker is a centralized repository for security-related information about Debian packages, including vulnerability reports, security advisories, and security updates. The tracker is designed to help users and administrators maintain the security of their Debian-based systems.

func (c *Client) GetIndexDebian(queryParameters ...IndexQueryParameters) (responseJSON *IndexDebianResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("debian"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDebianDsaResponse struct {
	Benchmark float64                                 `json:"_benchmark"`
	Meta      IndexMeta                               `json:"_meta"`
	Data      []client.AdvisoryDebianSecurityAdvisory `json:"data"`
}

// GetIndexDebianDsa -  Debian DSA (Debian Security Advisory) - `debian-dsa` index is a series of security advisories published by the Debian Project, a non-profit organization that develops and distributes the Debian operating system. These advisories provide information and guidance on security vulnerabilities and issues affecting Debian packages and software.

func (c *Client) GetIndexDebianDsa(queryParameters ...IndexQueryParameters) (responseJSON *IndexDebianDsaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("debian-dsa"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDellResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryDell `json:"data"`
}

// GetIndexDell -  Dell security advisories are official notifications released by Dell to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDell(queryParameters ...IndexQueryParameters) (responseJSON *IndexDellResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dell"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDeltaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryDeltaAdvisory `json:"data"`
}

// GetIndexDelta -  Delta Controls security bulletins are official notifications released by Delta Controls to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDelta(queryParameters ...IndexQueryParameters) (responseJSON *IndexDeltaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("delta"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDjangoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryDjango `json:"data"`
}

// GetIndexDjango -  Django security issues are official notifications released by the Django team to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDjango(queryParameters ...IndexQueryParameters) (responseJSON *IndexDjangoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("django"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDnnResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryDNN `json:"data"`
}

// GetIndexDnn -  DNN security advisories are official notifications released by DNN to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDnn(queryParameters ...IndexQueryParameters) (responseJSON *IndexDnnResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dnn"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDotcmsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryDotCMS `json:"data"`
}

// GetIndexDotcms -  dotCMS security advisories are official notifications released by dotCMS to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDotcms(queryParameters ...IndexQueryParameters) (responseJSON *IndexDotcmsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dotcms"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDragosResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryDragosAdvisory `json:"data"`
}

// GetIndexDragos -  Dragos security advisories are official notifications released by Dragos to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDragos(queryParameters ...IndexQueryParameters) (responseJSON *IndexDragosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dragos"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDraytekResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryDraytek `json:"data"`
}

// GetIndexDraytek -  DrayTek security advisories are official notifications released by DrayTek to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexDraytek(queryParameters ...IndexQueryParameters) (responseJSON *IndexDraytekResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("draytek"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexDrupalResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryDrupal `json:"data"`
}

// GetIndexDrupal -  Drupal security advisories are official notifications released by the Drupal Security Team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexDrupal(queryParameters ...IndexQueryParameters) (responseJSON *IndexDrupalResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("drupal"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEatonResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryEatonAdvisory `json:"data"`
}

// GetIndexEaton -  Eaton Security Advisories typically include detailed technical information about the vulnerability or issue, as well as recommendations for remediation and risk mitigation. They may also include severity ratings and CVSS scores to help organizations prioritize their response to potential security incidents. Eaton's security team works closely with customers and partners to identify and address security concerns, and is committed to providing timely and effective security advisories to help protect critical assets and data.

func (c *Client) GetIndexEaton(queryParameters ...IndexQueryParameters) (responseJSON *IndexEatonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("eaton"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexElasticResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryElastic `json:"data"`
}

// GetIndexElastic -  Elasticsearch security advisories are official notifications released by Elasticsearch to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexElastic(queryParameters ...IndexQueryParameters) (responseJSON *IndexElasticResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("elastic"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexElspecResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryElspec `json:"data"`
}

// GetIndexElspec -  Elspec security advisories are official notifications released by Elspec to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexElspec(queryParameters ...IndexQueryParameters) (responseJSON *IndexElspecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("elspec"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEmergingThreatsSnortResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryEmergingThreatsSnort `json:"data"`
}

// GetIndexEmergingThreatsSnort -  Proofpoint's Emerging Threats Snort Rules are snort rules that can be used to monitor network traffic for malicious activity.

func (c *Client) GetIndexEmergingThreatsSnort(queryParameters ...IndexQueryParameters) (responseJSON *IndexEmergingThreatsSnortResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("emerging-threats-snort"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEmersonResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryEmersonAdvisory `json:"data"`
}

// GetIndexEmerson -  Emerson Cyber Security Notifications are official alerts and notifications provided by Emerson, a global technology and engineering company. These notifications highlight emerging cyber threats, vulnerabilities, and security updates related to Emerson's automation and control systems. They provide critical information, recommendations, and patches to enhance the cybersecurity posture of industrial environments and protect critical infrastructure from potential cyberattacks.

func (c *Client) GetIndexEmerson(queryParameters ...IndexQueryParameters) (responseJSON *IndexEmersonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("emerson"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEndoflifeResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryEndOfLife `json:"data"`
}

// GetIndexEndoflife -  End-of-life (EOL) and support information is often hard to track, or very badly presented. This index documents EOL dates and support lifecycles for various products.
func (c *Client) GetIndexEndoflife(queryParameters ...IndexQueryParameters) (responseJSON *IndexEndoflifeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("endoflife"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEolResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryEOLReleaseData `json:"data"`
}

// GetIndexEol -  The VulnCheck EOL index contains a set of operating systems with associated end-of-life and long term support information.

func (c *Client) GetIndexEol(queryParameters ...IndexQueryParameters) (responseJSON *IndexEolResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("eol"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEolMicrosoftResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryEOLMicrosoft `json:"data"`
}

// GetIndexEolMicrosoft -  The Microsoft EOL data feed contains Microsoft product lifecycle data including release, retirement dates and support policies.
func (c *Client) GetIndexEolMicrosoft(queryParameters ...IndexQueryParameters) (responseJSON *IndexEolMicrosoftResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("eol-microsoft"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEpssResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiEPSSData `json:"data"`
}

// GetIndexEpss -  The Exploit Prediction Scoring System (EPSS) is a data-driven effort for estimating the probability that a software vulnerability will be exploited in the wild.

func (c *Client) GetIndexEpss(queryParameters ...IndexQueryParameters) (responseJSON *IndexEpssResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("epss"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEuvdResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryEUVD `json:"data"`
}

// GetIndexEuvd -  EUVD security advisories are official notifications released by the European Union to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexEuvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexEuvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("euvd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexExodusIntelResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryExodusIntel `json:"data"`
}

// GetIndexExodusIntel -  Exodus Intelligence advisories are official notifications released by Exodus Intelligence to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexExodusIntel(queryParameters ...IndexQueryParameters) (responseJSON *IndexExodusIntelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exodus-intel"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexExploitChainsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.ApiExploitChain `json:"data"`
}

// GetIndexExploitChains -  Exploit chains advisories are a type of security advisory that focus on the combination of multiple exploits or vulnerabilities that together create a more significant security risk. These advisories typically describe how an attacker could use multiple vulnerabilities in sequence to achieve a desired outcome, such as gaining unauthorized access to a system or stealing sensitive information.

func (c *Client) GetIndexExploitChains(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitChainsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploit-chains"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexExploitdbResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryExploitDBExploitv2 `json:"data"`
}

// GetIndexExploitdb -  The Exploit Database (ExploitDB) is an archive of public exploits curated by OffSec.

func (c *Client) GetIndexExploitdb(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitdbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploitdb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexExploitsResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.ApiExploitV3Result `json:"data"`
}

// GetIndexExploits -  VulnCheck Exploit Intelligence helps organizations track all of the world’s exploit proof-of-concept code, exploited in-the-wild information, and exploit metadata including timelines, to focus remediation resources on the right vulnerabilities.

func (c *Client) GetIndexExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploits"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexExploitsChangelogResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.ApiExploitsChangelog `json:"data"`
}

// GetIndexExploitsChangelog -  Provides a history of the changes made to an exploits record.
func (c *Client) GetIndexExploitsChangelog(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitsChangelogResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploits-changelog"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFSecureResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryFSecure `json:"data"`
}

// GetIndexFSecure -  F-Secure security advisories are official notifications released by F-Secure to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFSecure(queryParameters ...IndexQueryParameters) (responseJSON *IndexFSecureResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("f-secure"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexF5Response struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryF5 `json:"data"`
}

// GetIndexF5 -  F5 security advisories are official notifications released by F5 to address security vulnerabilities and updates in their software and hardware products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexF5(queryParameters ...IndexQueryParameters) (responseJSON *IndexF5Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("f5"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFanucResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryFanuc `json:"data"`
}

// GetIndexFanuc -  Fanuc security advisories are official notifications released by Fanuc to address security vulnerabilities and updates in their software and hardware products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFanuc(queryParameters ...IndexQueryParameters) (responseJSON *IndexFanucResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fanuc"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFastlyResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryFastly `json:"data"`
}

// GetIndexFastly -  Fastly security advisories are official notifications released by Fastly to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFastly(queryParameters ...IndexQueryParameters) (responseJSON *IndexFastlyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fastly"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFedoraResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUpdate `json:"data"`
}

// GetIndexFedora -  Fedora security advisories are official notifications released by Fedora to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFedora(queryParameters ...IndexQueryParameters) (responseJSON *IndexFedoraResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fedora"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFilecloudResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryFileCloud `json:"data"`
}

// GetIndexFilecloud -  FileCloud security advisories are official notifications released by FileCloud to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFilecloud(queryParameters ...IndexQueryParameters) (responseJSON *IndexFilecloudResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("filecloud"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFilezillaResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryFileZilla `json:"data"`
}

// GetIndexFilezilla -  FileZilla security advisories are official notifications released by FileZilla to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexFilezilla(queryParameters ...IndexQueryParameters) (responseJSON *IndexFilezillaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("filezilla"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFlattSecurityResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryFlattSecurity `json:"data"`
}

// GetIndexFlattSecurity -  Flatt Security advisories are official notifications released by Flatt Security to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFlattSecurity(queryParameters ...IndexQueryParameters) (responseJSON *IndexFlattSecurityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("flatt-security"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexForgerockResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryForgeRock `json:"data"`
}

// GetIndexForgerock -  ForgeRock security advisories are official notifications released by ForgeRock to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexForgerock(queryParameters ...IndexQueryParameters) (responseJSON *IndexForgerockResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("forgerock"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFortinetResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryFortinetAdvisory `json:"data"`
}

// GetIndexFortinet -  FortiGuard, by Fortinet, is a comprehensive and integrated security platform that offers threat intelligence, research, and protection against a wide range of cyber threats. It provides real-time updates on the latest threats and vulnerabilities, including malware, exploits, and botnets, enabling organizations to proactively defend their networks and systems. FortiGuard's threat intelligence and security services are a key component of Fortinet's security solutions, delivering advanced protection and continuous monitoring to safeguard against evolving cyber threats.

func (c *Client) GetIndexFortinet(queryParameters ...IndexQueryParameters) (responseJSON *IndexFortinetResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fortinet"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFortinetIpsResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryFortinetIPS `json:"data"`
}

// GetIndexFortinetIps -  The Fortinet Labs Threat Encyclopedia is a list of threats identified by Fortinet.
func (c *Client) GetIndexFortinetIps(queryParameters ...IndexQueryParameters) (responseJSON *IndexFortinetIpsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fortinet-ips"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFoxitResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryFoxit `json:"data"`
}

// GetIndexFoxit -  Foxit security bulletins are official notifications released by Foxit to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFoxit(queryParameters ...IndexQueryParameters) (responseJSON *IndexFoxitResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("foxit"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexFreebsdResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAdvisory `json:"data"`
}

// GetIndexFreebsd -  FreeBSD security advisories are official notifications released by the FreeBSD security team to address security vulnerabilities and updates in the open source FreeBSD operating system. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexFreebsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexFreebsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("freebsd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGallagherResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryGallagher `json:"data"`
}

// GetIndexGallagher -  Gallagher security advisories are official notifications released by Gallagher to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGallagher(queryParameters ...IndexQueryParameters) (responseJSON *IndexGallagherResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gallagher"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGcpResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryGCP `json:"data"`
}

// GetIndexGcp -  GCP security bulletins are official notifications released by Google Cloud to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGcp(queryParameters ...IndexQueryParameters) (responseJSON *IndexGcpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gcp"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGeGasResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryGEGas `json:"data"`
}

// GetIndexGeGas -  GE Gas product security advisories are official notifications released by the GE Gas Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGeGas(queryParameters ...IndexQueryParameters) (responseJSON *IndexGeGasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ge-gas"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGeHealthcareResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryGEHealthcareAdvisory `json:"data"`
}

// GetIndexGeHealthcare -  GE Healthcare Advisories are official communications issued by GE Healthcare, a global medical technology company, to provide information and guidance on potential security vulnerabilities and threats affecting GE Healthcare products and services.

func (c *Client) GetIndexGeHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *IndexGeHealthcareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ge-healthcare"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGemResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexGem -  Gem (Ruby) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexGem(queryParameters ...IndexQueryParameters) (responseJSON *IndexGemResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gem"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGenetecResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryGenetec `json:"data"`
}

// GetIndexGenetec -  Genetec security advisories are official notifications released by Genetec to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGenetec(queryParameters ...IndexQueryParameters) (responseJSON *IndexGenetecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("genetec"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGigabyteResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryGigabyte `json:"data"`
}

// GetIndexGigabyte -  Gigabyte security advisories are official notifications released by Gigabyte to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGigabyte(queryParameters ...IndexQueryParameters) (responseJSON *IndexGigabyteResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gigabyte"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGiteeExploitsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryGiteeExploit `json:"data"`
}

// GetIndexGiteeExploits -  | Exploits hosted on Gitee

func (c *Client) GetIndexGiteeExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexGiteeExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gitee-exploits"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGithubExploitsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryGitHubExploit `json:"data"`
}

// GetIndexGithubExploits -  | Exploits hosted on GitHub

func (c *Client) GetIndexGithubExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexGithubExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("github-exploits"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGithubSecurityAdvisoriesResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryGHAdvisoryJSONLean `json:"data"`
}

// GetIndexGithubSecurityAdvisories -  Github Security Advisories are official notifications released by Github to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexGithubSecurityAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexGithubSecurityAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("github-security-advisories"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGitlabAdvisoriesCommunityResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryGitlabAdvisory `json:"data"`
}

// GetIndexGitlabAdvisoriesCommunity -  GitLab Advisories Community is a group of security researchers and professionals who collaborate to identify and report security vulnerabilities and issues affecting the GitLab software development platform.

func (c *Client) GetIndexGitlabAdvisoriesCommunity(queryParameters ...IndexQueryParameters) (responseJSON *IndexGitlabAdvisoriesCommunityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gitlab-advisories-community"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGitlabExploitsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryGitLabExploit `json:"data"`
}

// GetIndexGitlabExploits -  | Exploits hosted on GitLab

func (c *Client) GetIndexGitlabExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexGitlabExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gitlab-exploits"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGnutlsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryGnuTLS `json:"data"`
}

// GetIndexGnutls -  GnuTLS security advisories are official notifications released by the GnuTLS open source project to address security vulnerabilities and updates in curl. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGnutls(queryParameters ...IndexQueryParameters) (responseJSON *IndexGnutlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gnutls"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGoVulndbResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryGoVulnJSON `json:"data"`
}

// GetIndexGoVulndb -  Data about new vulnerabilities come directly from Go package maintainers or sources such as MITRE and GitHub. Reports are curated by the Go Security team.
func (c *Client) GetIndexGoVulndb(queryParameters ...IndexQueryParameters) (responseJSON *IndexGoVulndbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("go-vulndb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGolangResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexGolang -  Golang packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexGolang(queryParameters ...IndexQueryParameters) (responseJSON *IndexGolangResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("golang"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGoogle0dayItwResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryITWExploit `json:"data"`
}

// GetIndexGoogle0dayItw -  Project Zero's In the Wild Exploits exploits list are curated by Google's Project Zero team and tracks zero day exploits found in the wild.

func (c *Client) GetIndexGoogle0dayItw(queryParameters ...IndexQueryParameters) (responseJSON *IndexGoogle0dayItwResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("google-0day-itw"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGoogleContainerOptimizedOsResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryContainerOS `json:"data"`
}

// GetIndexGoogleContainerOptimizedOs -  Container OS security advisories are official notifications released by Google to address security vulnerabilities and updates in the container optimized operating system. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGoogleContainerOptimizedOs(queryParameters ...IndexQueryParameters) (responseJSON *IndexGoogleContainerOptimizedOsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("google-container-optimized-os"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGrafanaResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryGrafana `json:"data"`
}

// GetIndexGrafana -  Grafana Labs security fixes are official notifications released by Grafana Labs to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexGrafana(queryParameters ...IndexQueryParameters) (responseJSON *IndexGrafanaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("grafana"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGreynoiseMetadataResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryGreyNoiseDetection `json:"data"`
}

// GetIndexGreynoiseMetadata -  GreyNoise Metadata Advisories are a type of security advisory that provides information about metadata associated with various IP addresses, domains, and other internet-connected devices.

func (c *Client) GetIndexGreynoiseMetadata(queryParameters ...IndexQueryParameters) (responseJSON *IndexGreynoiseMetadataResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("greynoise-metadata"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHackageResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexHackage -  Hackage (Haskell) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexHackage(queryParameters ...IndexQueryParameters) (responseJSON *IndexHackageResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hackage"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHarmonyosResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHarmonyOS `json:"data"`
}

// GetIndexHarmonyos -  HarmonyOS security updates are official notifications released by the HarmonyOS security team to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHarmonyos(queryParameters ...IndexQueryParameters) (responseJSON *IndexHarmonyosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("harmonyos"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHashicorpResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHashiCorp `json:"data"`
}

// GetIndexHashicorp -  HashiCorp security updates are official notifications released by HashiCorp to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHashicorp(queryParameters ...IndexQueryParameters) (responseJSON *IndexHashicorpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hashicorp"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHaskellSadbResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryHaskellSADBAdvisory `json:"data"`
}

// GetIndexHaskellSadb -  The Haskell Security Advisory Database is a repository of security advisories filed against packages published via Hackage.

func (c *Client) GetIndexHaskellSadb(queryParameters ...IndexQueryParameters) (responseJSON *IndexHaskellSadbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("haskell-sadb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHclResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryHCL `json:"data"`
}

// GetIndexHcl -  HCLSoftware security bulletins are official notifications released by the HCLSoftware Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHcl(queryParameters ...IndexQueryParameters) (responseJSON *IndexHclResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hcl"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHexResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexHex -  Hex (Erlang, Elixir) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexHex(queryParameters ...IndexQueryParameters) (responseJSON *IndexHexResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hex"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHikvisionResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHIKVision `json:"data"`
}

// GetIndexHikvision -  Hikvision security advisories are official notifications released by Hikvision to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHikvision(queryParameters ...IndexQueryParameters) (responseJSON *IndexHikvisionResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hikvision"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHillromResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryHillromAdvisory `json:"data"`
}

// GetIndexHillrom -  Hillrom Advisories are official notifications released by Hillrom, a leading global medical technology company. These advisories address security vulnerabilities and updates in Hillrom's medical devices and healthcare IT solutions. They provide critical information about potential risks, recommended actions, and available patches or updates to ensure the security and privacy of patient data and the proper functioning of Hillrom products.

func (c *Client) GetIndexHillrom(queryParameters ...IndexQueryParameters) (responseJSON *IndexHillromResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hillrom"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHitachiResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryHitachi `json:"data"`
}

// GetIndexHitachi -  Hitachi Software Vulnerability Information provides updates and notifications about security vulnerabilities and related software updates in Hitachi's software products. These notifications highlight potential risks, impacts, and recommended actions to mitigate vulnerabilities and protect systems from cyber threats.

func (c *Client) GetIndexHitachi(queryParameters ...IndexQueryParameters) (responseJSON *IndexHitachiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hitachi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHitachiEnergyResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryHitachiEnergy `json:"data"`
}

// GetIndexHitachiEnergy -  Hitachi Energy cybersecurity advisories and notifications are official notifications released by Hitachi Energy to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHitachiEnergy(queryParameters ...IndexQueryParameters) (responseJSON *IndexHitachiEnergyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hitachi-energy"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHkcertResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryHKCert `json:"data"`
}

// GetIndexHkcert -  Hong Kong CERT security bulletins are official notifications released by the Hong Kong CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexHkcert(queryParameters ...IndexQueryParameters) (responseJSON *IndexHkcertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hkcert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHoneywellResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHoneywell `json:"data"`
}

// GetIndexHoneywell -  Honeywell cyber security notifications are official notifications released by Honeywell to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHoneywell(queryParameters ...IndexQueryParameters) (responseJSON *IndexHoneywellResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("honeywell"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHpResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryHP `json:"data"`
}

// GetIndexHp -  HP security bulletins are official notifications released by HP to address security vulnerabilities and updates in their software and hardware products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHp(queryParameters ...IndexQueryParameters) (responseJSON *IndexHpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hp"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHpeResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryHPE `json:"data"`
}

// GetIndexHpe -  HPE security advisories are official notifications released by HP Enterprise to address security vulnerabilities and updates in their software and hardware products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexHpe(queryParameters ...IndexQueryParameters) (responseJSON *IndexHpeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hpe"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHuaweiEulerosResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryHuaweiEulerOS `json:"data"`
}

// GetIndexHuaweiEuleros -  OpenEuler Open Enterprise Operating System Security Advisories are official notifications released by the EulerOS security team to address security vulnerabilities and updates in the open enterprise EulerOS operating system. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHuaweiEuleros(queryParameters ...IndexQueryParameters) (responseJSON *IndexHuaweiEulerosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("huawei-euleros"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHuaweiIpsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHuaweiIPS `json:"data"`
}

// GetIndexHuaweiIps -  Huawei IPS Vulnerabilities are official notifications released by Huawei to address security vulnerabilities caught by Huawei's Intrusion Prevention System. These vulnerability notifications provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexHuaweiIps(queryParameters ...IndexQueryParameters) (responseJSON *IndexHuaweiIpsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("huawei-ips"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexHuaweiPsirtResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryHuawei `json:"data"`
}

// GetIndexHuaweiPsirt -  Huawei PSIRT seucrity bulletins are official notifications released by the Huawei Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexHuaweiPsirt(queryParameters ...IndexQueryParameters) (responseJSON *IndexHuaweiPsirtResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("huawei-psirt"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIavaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryIAVA `json:"data"`
}

// GetIndexIava -  Notifications that are generated when an Information Assurance vulnerability may result in an immediate and potentially severe threat to DoD systems and information; this alert requires corrective action because of the severity of the vulnerability risk.

func (c *Client) GetIndexIava(queryParameters ...IndexQueryParameters) (responseJSON *IndexIavaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("iava"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIbmResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryIBM `json:"data"`
}

// GetIndexIbm -  IBM security bulletins are official notifications released by IBM to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexIbm(queryParameters ...IndexQueryParameters) (responseJSON *IndexIbmResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ibm"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIdemiaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryIdemia `json:"data"`
}

// GetIndexIdemia -  Idemia product security vulnerabilities are official notifications released by the Idemia Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexIdemia(queryParameters ...IndexQueryParameters) (responseJSON *IndexIdemiaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("idemia"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIlAlertsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryIsraeliAlert `json:"data"`
}

// GetIndexIlAlerts -  Gov.il Security Alerts are official notifications issued by the Israeli government to provide timely information and updates on cybersecurity threats, vulnerabilities, and incidents. These alerts aim to raise awareness among government entities, critical infrastructure sectors, and the public about emerging cyber threats and provide recommended actions to mitigate risks.

func (c *Client) GetIndexIlAlerts(queryParameters ...IndexQueryParameters) (responseJSON *IndexIlAlertsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("il-alerts"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIlVulnerabilitiesResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryIsraeliVulnerability `json:"data"`
}

// GetIndexIlVulnerabilities -  Gov.il CVE Security Advisories are official notifications released by the Israeli government to address security vulnerabilities identified through the Common Vulnerabilities and Exposures (CVE) system. These advisories provide detailed information about specific vulnerabilities, their potential impact, and recommended actions to mitigate the risks.

func (c *Client) GetIndexIlVulnerabilities(queryParameters ...IndexQueryParameters) (responseJSON *IndexIlVulnerabilitiesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("il-vulnerabilities"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIncibeResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryIncibeAdvisory `json:"data"`
}

// GetIndexIncibe -  Incibe CERT early warnings are official notifications released by the  National Cybersecurity Institute of Spain (Incibe) to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexIncibe(queryParameters ...IndexQueryParameters) (responseJSON *IndexIncibeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("incibe"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexInitialAccessResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiInitialAccess `json:"data"`
}

// GetIndexInitialAccess -  The initial-access index contains data on Initial Access exploits. These exploits are typically the most high impact exploit published. These vulnerabilities, also sometimes referred to as Remote Code Execution (RCE) vulnerabilities, are remote in nature, and typically do not require credentials to exploit.

func (c *Client) GetIndexInitialAccess(queryParameters ...IndexQueryParameters) (responseJSON *IndexInitialAccessResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("initial-access"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexInitialAccessGitResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiInitialAccess `json:"data"`
}

// GetIndexInitialAccessGit -  This is a backup-only index for Initial Access detection artifacts hosted on git.vulncheck.com. This backup is only available to licensed subscribers of Initial Access Intelligence.

func (c *Client) GetIndexInitialAccessGit(queryParameters ...IndexQueryParameters) (responseJSON *IndexInitialAccessGitResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("initial-access-git"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIntelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryIntel `json:"data"`
}

// GetIndexIntel -  Intel Product Security Center advisories are official notifications released by Intel to address security vulnerabilities and updates in their software and hardware products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexIntel(queryParameters ...IndexQueryParameters) (responseJSON *IndexIntelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("intel"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel10dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

// GetIndexIpintel10d -  The 10-Day IP Intelligence index contains the IP address and geolocation of potentially vulnerable systems that may be targeted by initial access exploits as well as command and control (C2) attacker infrastructure.

func (c *Client) GetIndexIpintel10d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel10dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-10d"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel30dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

// GetIndexIpintel30d -  The 30-Day IP Intelligence index contains the IP address and geolocation of potentially vulnerable systems that may be targeted by initial access exploits as well as command and control (C2) attacker infrastructure.

func (c *Client) GetIndexIpintel30d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel30dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-30d"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel3dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

// GetIndexIpintel3d -  The 3-Day IP Intelligence index contains the IP address and geolocation of potentially vulnerable systems that may be targeted by initial access exploits as well as command and control (C2) attacker infrastructure.

func (c *Client) GetIndexIpintel3d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel3dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-3d"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel90dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

// GetIndexIpintel90d -  The 90-Day IP Intelligence index contains the IP address and geolocation of potentially vulnerable systems that may be targeted by initial access exploits as well as command and control (C2) attacker infrastructure.

func (c *Client) GetIndexIpintel90d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel90dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-90d"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIstioResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryIstio `json:"data"`
}

// GetIndexIstio -  Istio security bulletins are official notifications released by the open source Istio project to address security vulnerabilities and updates in the open source Istio project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexIstio(queryParameters ...IndexQueryParameters) (responseJSON *IndexIstioResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("istio"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIvantiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryIvanti `json:"data"`
}

// GetIndexIvanti -  Ivanti security updates are official notifications released by Ivanti to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexIvanti(queryParameters ...IndexQueryParameters) (responseJSON *IndexIvantiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ivanti"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIvantiRssResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryIvantiRSS `json:"data"`
}

// GetIndexIvantiRss -  Ivanti security advisories are official notifications released by Ivanti to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexIvantiRss(queryParameters ...IndexQueryParameters) (responseJSON *IndexIvantiRssResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ivanti-rss"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJenkinsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryJenkins `json:"data"`
}

// GetIndexJenkins -  Jenkins security advisories are official notifications released by Jenkins to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexJenkins(queryParameters ...IndexQueryParameters) (responseJSON *IndexJenkinsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jenkins"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJetbrainsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryJetBrains `json:"data"`
}

// GetIndexJetbrains -  JetBrains security issues are official notifications released by JetBrains to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexJetbrains(queryParameters ...IndexQueryParameters) (responseJSON *IndexJetbrainsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jetbrains"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJfrogResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryJFrog `json:"data"`
}

// GetIndexJfrog -  JFrog security advisories are official notifications released by JFrog to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexJfrog(queryParameters ...IndexQueryParameters) (responseJSON *IndexJfrogResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jfrog"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJnjResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryJNJAdvisory `json:"data"`
}

// GetIndexJnj -  Johnson & Johnson's Vulnerability Disclosure Reporting is a process through which individuals or security researchers can responsibly report potential vulnerabilities in Johnson & Johnson's products, services, or systems.

func (c *Client) GetIndexJnj(queryParameters ...IndexQueryParameters) (responseJSON *IndexJnjResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jnj"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJvnResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryJVN `json:"data"`
}

// GetIndexJvn -  JVN stands for "the Japan Vulnerability Notes." It is a vulnerability information portal site designed to help ensure Internet security by providing vulnerability information and their solutions for software products used in Japan. JVN is operated jointly by the JPCERT Coordination Center and the Information-technology Promotion Agency (IPA).
func (c *Client) GetIndexJvn(queryParameters ...IndexQueryParameters) (responseJSON *IndexJvnResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jvn"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJvndbResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryJVNAdvisoryItem `json:"data"`
}

// GetIndexJvndb -  Japan vulnerability notes are official notifications released by the Japan CERT (JPCERT) to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexJvndb(queryParameters ...IndexQueryParameters) (responseJSON *IndexJvndbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jvndb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexKasperskyIcsCertResponse struct {
	Benchmark float64                                   `json:"_benchmark"`
	Meta      IndexMeta                                 `json:"_meta"`
	Data      []client.AdvisoryKasperskyICSCERTAdvisory `json:"data"`
}

// GetIndexKasperskyIcsCert -  Kaspersky ICS CERT (Industrial Control Systems Computer Emergency Response Team) is a specialized unit within Kaspersky that focuses on cybersecurity for industrial control systems.

func (c *Client) GetIndexKasperskyIcsCert(queryParameters ...IndexQueryParameters) (responseJSON *IndexKasperskyIcsCertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("kaspersky-ics-cert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexKorelogicResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryKoreLogic `json:"data"`
}

// GetIndexKorelogic -  KoreLogic vulnerability research and advisories are official notifications released by KoreLogic to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexKorelogic(queryParameters ...IndexQueryParameters) (responseJSON *IndexKorelogicResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("korelogic"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexKrcertSecurityNoticesResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryKRCertAdvisory `json:"data"`
}

// GetIndexKrcertSecurityNotices -  KR-CERT (Korea Internet & Security Agency Computer Emergency Response Team) Security Notices are official notifications issued by KR-CERT, the national computer emergency response team of South Korea.

func (c *Client) GetIndexKrcertSecurityNotices(queryParameters ...IndexQueryParameters) (responseJSON *IndexKrcertSecurityNoticesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("krcert-security-notices"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexKrcertVulnerabilitiesResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryKRCertAdvisory `json:"data"`
}

// GetIndexKrcertVulnerabilities -  KR-CERT (Korea Internet & Security Agency Computer Emergency Response Team) provides valuable information on vulnerabilities that affect the South Korean cyberspace.

func (c *Client) GetIndexKrcertVulnerabilities(queryParameters ...IndexQueryParameters) (responseJSON *IndexKrcertVulnerabilitiesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("krcert-vulnerabilities"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexKubernetesResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryK8S `json:"data"`
}

// GetIndexKubernetes -  Kubernetes security issues are official notifications released by the Kubernetes Security Response Committee to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexKubernetes(queryParameters ...IndexQueryParameters) (responseJSON *IndexKubernetesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("kubernetes"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexLenovoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryLenovo `json:"data"`
}

// GetIndexLenovo -  Lenovo product security advisories are official notifications released by the Lenovo Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexLenovo(queryParameters ...IndexQueryParameters) (responseJSON *IndexLenovoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("lenovo"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexLexmarkResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryLexmarkAdvisory `json:"data"`
}

// GetIndexLexmark -  Lexmark security advisories are official notifications released by Lexmark to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexLexmark(queryParameters ...IndexQueryParameters) (responseJSON *IndexLexmarkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("lexmark"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexLgResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryLG `json:"data"`
}

// GetIndexLg -  LG security bulletins are official notifications released by LG to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexLg(queryParameters ...IndexQueryParameters) (responseJSON *IndexLgResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("lg"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexLibreOfficeResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryLibreOffice `json:"data"`
}

// GetIndexLibreOffice -  Libre Office security advisories are official notifications released by the open source Libre Office project to address security vulnerabilities and updates in the open source Libre Office project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexLibreOffice(queryParameters ...IndexQueryParameters) (responseJSON *IndexLibreOfficeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("libre-office"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexLinuxResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryLinux `json:"data"`
}

// GetIndexLinux -  Linux kernel security advisories are official notifications released by the Linux security team to address security vulnerabilities and updates in the open source Linux operating system. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexLinux(queryParameters ...IndexQueryParameters) (responseJSON *IndexLinuxResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("linux"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMFilesResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMFiles `json:"data"`
}

// GetIndexMFiles -  M-Files security advisories are official notifications released by M-Files to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMFiles(queryParameters ...IndexQueryParameters) (responseJSON *IndexMFilesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("m-files"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMacertResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMACert `json:"data"`
}

// GetIndexMacert -  Moroccan CERT security bulletins are official notifications released by the Moroccan CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexMacert(queryParameters ...IndexQueryParameters) (responseJSON *IndexMacertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("macert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexManageengineResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryManageEngineAdvisory `json:"data"`
}

// GetIndexManageengine -  ManageEngine security updates are official notifications released by the ManageEngine Security Response Center to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexManageengine(queryParameters ...IndexQueryParameters) (responseJSON *IndexManageengineResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("manageengine"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMavenResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexMaven -  Maven (Java) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexMaven(queryParameters ...IndexQueryParameters) (responseJSON *IndexMavenResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("maven"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMbedTlsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMbedTLS `json:"data"`
}

// GetIndexMbedTls -  Mbed TLS security advisories are official notifications released by the open source Mbed TLS project to address security vulnerabilities and updates in the open source Mbed TLS project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMbedTls(queryParameters ...IndexQueryParameters) (responseJSON *IndexMbedTlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mbed-tls"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMcafeeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMcAfee `json:"data"`
}

// GetIndexMcafee -  McAfee security advisories are official notifications released by McAfee to address security vulnerabilities and updates in their software and hardware products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexMcafee(queryParameters ...IndexQueryParameters) (responseJSON *IndexMcafeeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mcafee"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMediatekResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMediatek `json:"data"`
}

// GetIndexMediatek -  MediaTek security advisories are official notifications released by MediaTek to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMediatek(queryParameters ...IndexQueryParameters) (responseJSON *IndexMediatekResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mediatek"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMedtronicResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryMedtronicAdvisory `json:"data"`
}

// GetIndexMedtronic -  Medtronic security bulletins are official notifications released by Medtronic to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMedtronic(queryParameters ...IndexQueryParameters) (responseJSON *IndexMedtronicResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("medtronic"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMendixResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMendix `json:"data"`
}

// GetIndexMendix -  Mendix security advisories are official notifications released by the Siemens ProductCERT Team to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMendix(queryParameters ...IndexQueryParameters) (responseJSON *IndexMendixResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mendix"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMetasploitResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryMetasploitExploit `json:"data"`
}

// GetIndexMetasploit -  Metasploit Modules is a list of modules that can be utilized via the metasploit framework for pentesting.

func (c *Client) GetIndexMetasploit(queryParameters ...IndexQueryParameters) (responseJSON *IndexMetasploitResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("metasploit"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMicrosoftCvrfResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryMicrosoftCVRF `json:"data"`
}

// GetIndexMicrosoftCvrf -  Microsoft Security Updates are official notifications released by the Microsoft Security Response Center (MSRC) to address security vulnerabilities and updates for Microsoft. These security updates provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMicrosoftCvrf(queryParameters ...IndexQueryParameters) (responseJSON *IndexMicrosoftCvrfResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("microsoft-cvrf"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMicrosoftKbResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryMicrosoftKb `json:"data"`
}

// GetIndexMicrosoftKb -  This data is a reformatted view of microsoft-cvrf showing each CVE and its list of KBs.

func (c *Client) GetIndexMicrosoftKb(queryParameters ...IndexQueryParameters) (responseJSON *IndexMicrosoftKbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("microsoft-kb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMikrotikResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMikrotik `json:"data"`
}

// GetIndexMikrotik -  MikroTik security bulletins are official notifications released by MikroTik to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMikrotik(queryParameters ...IndexQueryParameters) (responseJSON *IndexMikrotikResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mikrotik"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMindrayResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMindray `json:"data"`
}

// GetIndexMindray -  Mindray cybersecurity advisories are official notifications released by Mindray to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMindray(queryParameters ...IndexQueryParameters) (responseJSON *IndexMindrayResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mindray"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMispThreatActorsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryMispValue `json:"data"`
}

// GetIndexMispThreatActors -  MISP Threat Actors is an open source list of known threat actors for the MISP (Malware Information Sharing Program) Open Source Threat Intelligence Sharing Platform.

func (c *Client) GetIndexMispThreatActors(queryParameters ...IndexQueryParameters) (responseJSON *IndexMispThreatActorsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("misp-threat-actors"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMitelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryMitel `json:"data"`
}

// GetIndexMitel -  Mitel security advisories are official notifications released by Mitel to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMitel(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitel"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMitreAttackCveResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiMitreAttackToCVE `json:"data"`
}

// GetIndexMitreAttackCve -  Provides a map between certain MITRE ATT&CK technique IDs and applicable CVEs.

func (c *Client) GetIndexMitreAttackCve(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitreAttackCveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitre-attack-cve"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMitreCvelistV5Response struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryMitreCVEListV5 `json:"data"`
}

// GetIndexMitreCvelistV5 -  MITRE CVE is a collection of publicly disclosed cybersecurity vulnerabilities by NIST that aims to identify, define and catalog publicly disclosed cybersecurity vulnerabilities.

func (c *Client) GetIndexMitreCvelistV5(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitreCvelistV5Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitre-cvelist-v5"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMitsubishiElectricResponse struct {
	Benchmark float64                                     `json:"_benchmark"`
	Meta      IndexMeta                                   `json:"_meta"`
	Data      []client.AdvisoryMitsubishiElectricAdvisory `json:"data"`
}

// GetIndexMitsubishiElectric -  Mitsubishi Electric Vulnerabilities are official notifications released by the Mitsubishi Electric PSIRT (Product Security Incident Response Team) to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMitsubishiElectric(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitsubishiElectricResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitsubishi-electric"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMongodbResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMongoDB `json:"data"`
}

// GetIndexMongodb -  MongoDB security alerts are official notifications released by MongoDB to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMongodb(queryParameters ...IndexQueryParameters) (responseJSON *IndexMongodbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mongodb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMoxaResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryMoxaAdvisory `json:"data"`
}

// GetIndexMoxa -  Moxa security advisories are official notifications released by the Moxa Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMoxa(queryParameters ...IndexQueryParameters) (responseJSON *IndexMoxaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("moxa"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMozillaResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryMozillaAdvisory `json:"data"`
}

// GetIndexMozilla -  Mozilla security advisories are official notifications released by the Mozilla Foundation to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexMozilla(queryParameters ...IndexQueryParameters) (responseJSON *IndexMozillaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mozilla"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNaverResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryNaver `json:"data"`
}

// GetIndexNaver -  Naver security advisories are official notifications released by the Naver Security Team to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNaver(queryParameters ...IndexQueryParameters) (responseJSON *IndexNaverResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("naver"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNcscResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryNCSC `json:"data"`
}

// GetIndexNcsc -  Nationaal Cyber Security Centrum advisories are official notifications released by the Nationaal Cyber Security Centrum to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexNcsc(queryParameters ...IndexQueryParameters) (responseJSON *IndexNcscResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ncsc"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNcscCvesResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNCSCCVE `json:"data"`
}

// GetIndexNcscCves -  Nationaal Cyber Security Centrum cves are official notifications released by the Nationaal Cyber Security Centrum to address security vulnerabilities and updates. These cves provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexNcscCves(queryParameters ...IndexQueryParameters) (responseJSON *IndexNcscCvesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ncsc-cves"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNecResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNEC `json:"data"`
}

// GetIndexNec -  NEC security information notices are official notifications released by NEC to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNec(queryParameters ...IndexQueryParameters) (responseJSON *IndexNecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nec"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNetappResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNetApp `json:"data"`
}

// GetIndexNetapp -  NetApp Security Advisories are official notifications released by the NetApp PSIRT (Product Security Incident Response Team) to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNetapp(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetappResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netapp"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNetatalkResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryNetatalk `json:"data"`
}

// GetIndexNetatalk -  Netatalk security advisories are official notifications released by the Netatalk team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexNetatalk(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetatalkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netatalk"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNetgateResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNetgate `json:"data"`
}

// GetIndexNetgate -  Netgate security advisories are official notifications released by Netgate to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNetgate(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetgateResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netgate"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNetgearResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNetgear `json:"data"`
}

// GetIndexNetgear -  NETGEAR Security Advisories are official notifications released by NETGEAR to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNetgear(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetgearResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netgear"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNetskopeResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryNetskope `json:"data"`
}

// GetIndexNetskope -  Netskope security advisories are official notifications released by Netskope to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNetskope(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetskopeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netskope"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNginxResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryNginxAdvisory `json:"data"`
}

// GetIndexNginx -  Nginx security advisories are official notifications released by F5 to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNginx(queryParameters ...IndexQueryParameters) (responseJSON *IndexNginxResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nginx"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNhsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNHS `json:"data"`
}

// GetIndexNhs -  NHS cyber alerts are official notifications released by NHS Digital to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNhs(queryParameters ...IndexQueryParameters) (responseJSON *IndexNhsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nhs"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNiResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryNI `json:"data"`
}

// GetIndexNi -  National Instruments (NI) security updates are official notifications released by NI to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNi(queryParameters ...IndexQueryParameters) (responseJSON *IndexNiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ni"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNistNvdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiCveItems `json:"data"`
}

// GetIndexNistNvd -  NIST NVD (National Institute of Standards and Technology National Vulnerability Database) version 1.0 is an early release of a comprehensive repository of vulnerability information and security-related data. It serves as a valuable resource for cybersecurity professionals, researchers, and organizations by providing detailed information on known software vulnerabilities, including their severity, impact, and associated references. NVD version 1.0 offers a structured format for accessing and analyzing vulnerability data, aiding in the identification and mitigation of security risks across various software and hardware products.

func (c *Client) GetIndexNistNvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNistNvd2Response struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiNVD20CVE `json:"data"`
}

// GetIndexNistNvd2 -  The National Institute of Standards and Technology (NIST) National Vulnerability Database (NVD) v2.0 is a comprehensive repository of security vulnerability data, including Common Vulnerabilities and Exposures (CVEs). It provides a variety of information on CVEs, such as their severity, impact, and remediation strategies. NVD v2.0 also provides a Common Vulnerability Scoring System (CVSS) v2.0 calculator, which allows users to calculate the severity of a CVE based on its specific characteristics.

func (c *Client) GetIndexNistNvd2(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvd2Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd2"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNistNvd2CpematchResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiNVD20CPEMatch `json:"data"`
}

// GetIndexNistNvd2Cpematch -  NIST NVD 2.0 CPE Match Advisories are a type of security advisory that provides information about Common Platform Enumeration (CPE) matches associated with vulnerabilities in the National Vulnerability Database (NVD) 2.0. CPEs are standardized identifiers for software applications, operating systems, and other IT systems, and are used to help organizations identify and track vulnerabilities and other security issues. NIST NVD 2.0 CPE Match Advisories provide information about the CPEs associated with specific vulnerabilities listed in the NVD 2.0. This information can help organizations better understand the scope and potential impact of a given vulnerability, and to take appropriate action to mitigate the associated risks. NIST NVD 2.0 CPE Match Advisories may also include information about known exploits or other factors that may increase the severity of a given vulnerability. By leveraging the information provided by NIST NVD 2.0 CPE Match Advisories, organizations can gain a deeper understanding of potential security risks and vulnerabilities, and develop more effective strategies for mitigating those risks. The advisories can also help organizations to prioritize their response to potential security incidents, and to ensure that critical systems and applications are appropriately secured and protected against advanced and persistent threats. Overall, NIST NVD 2.0 CPE Match Advisories are an important tool for organizations looking to maintain the security and integrity of their networks and systems.

func (c *Client) GetIndexNistNvd2Cpematch(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvd2CpematchResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd2-cpematch"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNistNvd2SourcesResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryNVD20Source `json:"data"`
}

// GetIndexNistNvd2Sources -  This Index contains information about CNAs, such as email addresses, official names, UUIDs used in NVD records. This allows us to lookup the UUIDs in NVD records and retrieve CNA names.
func (c *Client) GetIndexNistNvd2Sources(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvd2SourcesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd2-sources"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNodeSecurityResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryNodeSecurity `json:"data"`
}

// GetIndexNodeSecurity -  Node.js security working group advisories are official notifications released by the Node.js Security Working Group to address security vulnerabilities and updates in the node and npm software ecosystems. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNodeSecurity(queryParameters ...IndexQueryParameters) (responseJSON *IndexNodeSecurityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("node-security"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNodejsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNodeJS `json:"data"`
}

// GetIndexNodejs -  NodeJS security release notices are official notifications released by NodeJS to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexNodejs(queryParameters ...IndexQueryParameters) (responseJSON *IndexNodejsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nodejs"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNokiaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryNokia `json:"data"`
}

// GetIndexNokia -  Nokia product security advisories are official notifications released by the Nokia Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNokia(queryParameters ...IndexQueryParameters) (responseJSON *IndexNokiaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nokia"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNotepadplusplusResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryNotePadPlusPlus `json:"data"`
}

// GetIndexNotepadplusplus -  Notepad++ security advisories are official notifications released by the Notepad++ team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexNotepadplusplus(queryParameters ...IndexQueryParameters) (responseJSON *IndexNotepadplusplusResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("notepadplusplus"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNozomiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNozomi `json:"data"`
}

// GetIndexNozomi -  Nozomi Networks security advisories are official notifications released by the Nozomi Networks Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNozomi(queryParameters ...IndexQueryParameters) (responseJSON *IndexNozomiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nozomi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNpmResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexNpm -  NPM (Javascript, Typescript) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexNpm(queryParameters ...IndexQueryParameters) (responseJSON *IndexNpmResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("npm"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNtpResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNTP `json:"data"`
}

// GetIndexNtp -  NTP security issues are official notifications released by the NTP project to address security vulnerabilities and updates in the open source NTP project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNtp(queryParameters ...IndexQueryParameters) (responseJSON *IndexNtpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ntp"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNugetResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexNuget -  NuGet (.NET) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexNuget(queryParameters ...IndexQueryParameters) (responseJSON *IndexNugetResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nuget"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNvidiaResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisorySecurityBulletin `json:"data"`
}

// GetIndexNvidia -  NVIDIA security bulletins are official notifications released by NVIDIA to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexNvidia(queryParameters ...IndexQueryParameters) (responseJSON *IndexNvidiaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nvidia"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNzAdvisoriesResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryNZAdvisory `json:"data"`
}

// GetIndexNzAdvisories -  CERT NZ security advisories are official notifications released by the New Zealand CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexNzAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexNzAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nz-advisories"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOctopusDeployResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOctopusDeploy `json:"data"`
}

// GetIndexOctopusDeploy -  Octopus Deploy security advisories are official notifications released by Octopus Deploy to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOctopusDeploy(queryParameters ...IndexQueryParameters) (responseJSON *IndexOctopusDeployResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("octopus-deploy"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOktaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOkta `json:"data"`
}

// GetIndexOkta -  Okta security advisories are official notifications released by Okta to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOkta(queryParameters ...IndexQueryParameters) (responseJSON *IndexOktaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("okta"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOmronResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryOmron `json:"data"`
}

// GetIndexOmron -  Omron vulnerability advisories are official notifications released by Omron to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOmron(queryParameters ...IndexQueryParameters) (responseJSON *IndexOmronResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("omron"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOneEResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOneE `json:"data"`
}

// GetIndexOneE -  1E published product vulnerabilities are official notifications released by 1E to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOneE(queryParameters ...IndexQueryParameters) (responseJSON *IndexOneEResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("one-e"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpamResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexOpam -  opam (OCaml) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexOpam(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpamResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("opam"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpenCvdbResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryOpenCVDB `json:"data"`
}

// GetIndexOpenCvdb -  The Open Cloud Vulnerability & Security Issue Database are official notifications released to address security vulnerabilities and updates in all publicly known cloud vulnerabilities and CSP security issues. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOpenCvdb(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenCvdbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("open-cvdb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpenbsdResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenBSD `json:"data"`
}

// GetIndexOpenbsd -  OpenBSD security advisories are official notifications released by the OpenBSD security team to address security vulnerabilities and updates in the open source OpenBSD operating system. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOpenbsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenbsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openbsd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpenjdkResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenJDK `json:"data"`
}

// GetIndexOpenjdk -  OpenJDK security advisories are official notifications released by the OpenJDK team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOpenjdk(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenjdkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openjdk"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpensshResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenSSH `json:"data"`
}

// GetIndexOpenssh -  OpenSSH security advisories are official notifications released by the OpenSSH security team to address security vulnerabilities and updates in the open source OpenSSH project. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOpenssh(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpensshResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openssh"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpensslSecadvResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOpenSSLSecAdv `json:"data"`
}

// GetIndexOpensslSecadv -  OpenSSL Security Advisories are official communications issued by the OpenSSL project, an open-source software library that provides cryptographic functions to protect communications over computer networks. These advisories are designed to provide information and guidance on potential security vulnerabilities and threats affecting OpenSSL software. OpenSSL Security Advisories typically include technical details about the vulnerability or issue, as well as recommended remediation and risk mitigation steps. They may also include severity ratings and CVSS scores to help organizations prioritize their response to potential security incidents. The OpenSSL security team works closely with the community to identify and address security concerns, and is committed to providing timely and effective security advisories to help protect user data and sensitive information. OpenSSL Security Advisories cover a wide range of topics, including vulnerabilities related to key management, cryptographic weaknesses, and protocol issues. By providing regular updates and guidance on potential security threats, OpenSSL helps to ensure the ongoing security and reliability of its software for its users. Additionally, OpenSSL encourages open and transparent collaboration with the community to help identify and address potential security concerns, making it an important component of secure communications infrastructure.

func (c *Client) GetIndexOpensslSecadv(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpensslSecadvResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openssl-secadv"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpenstackResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryOpenStack `json:"data"`
}

// GetIndexOpenstack -  OpenStack security advisories are official notifications released by the open source OpenStack project to address security vulnerabilities and updates in the open source OpenStack project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOpenstack(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenstackResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openstack"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOpenwrtResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryWRT `json:"data"`
}

// GetIndexOpenwrt -  OpenWRT security advisories are official notifications released by the OpenWRT team to address security vulnerabilities and updates in the open source OpenWRT operating system. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOpenwrt(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenwrtResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openwrt"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOracleResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMetaData `json:"data"`
}

// GetIndexOracle -  Oracle security advisories are official notifications released by Oracle to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems. category: Product Security Advisories

func (c *Client) GetIndexOracle(queryParameters ...IndexQueryParameters) (responseJSON *IndexOracleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("oracle"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOracleCpuResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryOracleCPU `json:"data"`
}

// GetIndexOracleCpu -  Oracle Critical Patch Updates provide security patches for supported Oracle on-premises products.

func (c *Client) GetIndexOracleCpu(queryParameters ...IndexQueryParameters) (responseJSON *IndexOracleCpuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("oracle-cpu"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOracleCpuCsafResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOracleCPUCSAF `json:"data"`
}

// GetIndexOracleCpuCsaf -  Oracle Critical Patch Updates provide security patches for supported Oracle on-premises products. These CPUs are released as CSAF on a quarterly basis.

func (c *Client) GetIndexOracleCpuCsaf(queryParameters ...IndexQueryParameters) (responseJSON *IndexOracleCpuCsafResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("oracle-cpu-csaf"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOsvResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryOSV `json:"data"`
}

// GetIndexOsv -  The Open Source Vulnerabilities Database is n open, precise, and distributed approach to producing and consuming vulnerability information for open source.

func (c *Client) GetIndexOsv(queryParameters ...IndexQueryParameters) (responseJSON *IndexOsvResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("osv"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOtrsResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOTRS `json:"data"`
}

// GetIndexOtrs -  OTRS security advisories are official notifications released by OTRS to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOtrs(queryParameters ...IndexQueryParameters) (responseJSON *IndexOtrsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("otrs"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexOwncloudResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryOwnCloud `json:"data"`
}

// GetIndexOwncloud -  OwnCloud security advisories are official notifications released by OwnCloud to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexOwncloud(queryParameters ...IndexQueryParameters) (responseJSON *IndexOwncloudResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("owncloud"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPacketstormResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryPacketstormExploit `json:"data"`
}

// GetIndexPacketstorm -  Packetstorm exploits is a list curated by the Packetstorm team that holds a quarter century of exploits.
func (c *Client) GetIndexPacketstorm(queryParameters ...IndexQueryParameters) (responseJSON *IndexPacketstormResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("packetstorm"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPalantirResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPalantir `json:"data"`
}

// GetIndexPalantir -  Palantir security bulletins are official notifications released by Palantir to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPalantir(queryParameters ...IndexQueryParameters) (responseJSON *IndexPalantirResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("palantir"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPaloAltoResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryPaloAltoAdvisory `json:"data"`
}

// GetIndexPaloAlto -  Palo Alto Networks Security Advisories are official notifications released by the Palo Alto Networks Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPaloAlto(queryParameters ...IndexQueryParameters) (responseJSON *IndexPaloAltoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("palo-alto"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPanasonicResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryPanasonic `json:"data"`
}

// GetIndexPanasonic -  The Panasonic vulnerability advisory list are official notifications released by the Panasonic Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPanasonic(queryParameters ...IndexQueryParameters) (responseJSON *IndexPanasonicResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("panasonic"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPapercutResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPaperCut `json:"data"`
}

// GetIndexPapercut -  PaperCut security vulnerabilities are official notifications released by PaperCut to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPapercut(queryParameters ...IndexQueryParameters) (responseJSON *IndexPapercutResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("papercut"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPegaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryPega `json:"data"`
}

// GetIndexPega -  Pega security bulletins are official notifications released by Pega to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPega(queryParameters ...IndexQueryParameters) (responseJSON *IndexPegaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pega"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPhilipsResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryPhilipsAdvisory `json:"data"`
}

// GetIndexPhilips -  Philips security advisories are official notifications released by Philips to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPhilips(queryParameters ...IndexQueryParameters) (responseJSON *IndexPhilipsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("philips"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPhoenixContactResponse struct {
	Benchmark float64                                 `json:"_benchmark"`
	Meta      IndexMeta                               `json:"_meta"`
	Data      []client.AdvisoryPhoenixContactAdvisory `json:"data"`
}

// GetIndexPhoenixContact -  Phoenix Contact security advisories are official notifications released by the Phoenix Contact Product Security Incident Response Team to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPhoenixContact(queryParameters ...IndexQueryParameters) (responseJSON *IndexPhoenixContactResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("phoenix-contact"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPhpMyAdminResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryPHPMyAdmin `json:"data"`
}

// GetIndexPhpMyAdmin -  phpMyAdmin security advisories are official notifications released by phpMyAdmin to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPhpMyAdmin(queryParameters ...IndexQueryParameters) (responseJSON *IndexPhpMyAdminResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("php-my-admin"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPostgressqlResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryPostgresSQL `json:"data"`
}

// GetIndexPostgressql -  PostgresSQL security vulnerabilities are official notifications released by the open source PostgresSQL project to address security vulnerabilities and updates in the open source PostgresSQL project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPostgressql(queryParameters ...IndexQueryParameters) (responseJSON *IndexPostgressqlResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("postgressql"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPowerdnsResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPowerDNS `json:"data"`
}

// GetIndexPowerdns -  PowerDNS security advisories are official notifications released by PowerDNS to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPowerdns(queryParameters ...IndexQueryParameters) (responseJSON *IndexPowerdnsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("powerdns"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexProgressResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryProgress `json:"data"`
}

// GetIndexProgress -  Progress product alert bulletins are official notifications released by Progress to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexProgress(queryParameters ...IndexQueryParameters) (responseJSON *IndexProgressResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("progress"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexProofpointResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryProofpoint `json:"data"`
}

// GetIndexProofpoint -  Proofpoint security advisories are official notifications released by Proofpoint to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexProofpoint(queryParameters ...IndexQueryParameters) (responseJSON *IndexProofpointResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("proofpoint"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPtcResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryPTC `json:"data"`
}

// GetIndexPtc -  PTC Security Advisories are official notifications released by PTC to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexPtc(queryParameters ...IndexQueryParameters) (responseJSON *IndexPtcResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ptc"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPubResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexPub -  Pub is a package manager for Dart and Flutter apps

func (c *Client) GetIndexPub(queryParameters ...IndexQueryParameters) (responseJSON *IndexPubResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pub"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPureStorageResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryPureStorage `json:"data"`
}

// GetIndexPureStorage -  Pure Storage security bulletins are official notifications released by Pure Storage to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexPureStorage(queryParameters ...IndexQueryParameters) (responseJSON *IndexPureStorageResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pure-storage"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPypaAdvisoriesResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryPyPAAdvisory `json:"data"`
}

// GetIndexPypaAdvisories -  The Python Package Advisories index holds community maintained collection of security advisories for PyPI packages.

func (c *Client) GetIndexPypaAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexPypaAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pypa-advisories"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexPypiResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexPypi -  PyPI (Python) packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexPypi(queryParameters ...IndexQueryParameters) (responseJSON *IndexPypiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pypi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexQnapResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryQNAPAdvisory `json:"data"`
}

// GetIndexQnap -  QNAP security advisories are official notifications released by QNAP to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexQnap(queryParameters ...IndexQueryParameters) (responseJSON *IndexQnapResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qnap"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexQualcommResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryQualcomm `json:"data"`
}

// GetIndexQualcomm -  Qualcomm security bulletins are official notifications released by Qualcomm Technologies to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexQualcomm(queryParameters ...IndexQueryParameters) (responseJSON *IndexQualcommResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qualcomm"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexQualysResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryQualys `json:"data"`
}

// GetIndexQualys -  Qualys security advisories are official notifications released by Qualys to address software security flaws found by Qualys and can include proof of concept exploit code.

func (c *Client) GetIndexQualys(queryParameters ...IndexQueryParameters) (responseJSON *IndexQualysResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qualys"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexQubesQsbResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryQSB `json:"data"`
}

// GetIndexQubesQsb -  Qubes Security Bulletins are official notifications released by QubesOS to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexQubesQsb(queryParameters ...IndexQueryParameters) (responseJSON *IndexQubesQsbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qubes-qsb"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRansomwareResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryRansomwareExploit `json:"data"`
}

// GetIndexRansomware -  The VulnCheck Ransomware index contains data related to various ransomware. The index contains listings of ransomware groups and citations for the CVE they have been known to use.

func (c *Client) GetIndexRansomware(queryParameters ...IndexQueryParameters) (responseJSON *IndexRansomwareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ransomware"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRedhatResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryRedhatCVE `json:"data"`
}

// GetIndexRedhat -  Red Hat Security Advisories, commonly referred to as RHSA, are official notifications and updates provided by Red Hat, Inc., a leading provider of open-source solutions and enterprise Linux distributions. These advisories are a critical part of Red Hat's commitment to ensuring the security of their products and services.

func (c *Client) GetIndexRedhat(queryParameters ...IndexQueryParameters) (responseJSON *IndexRedhatResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("redhat"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRenesasResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryRenesas `json:"data"`
}

// GetIndexRenesas -  Renesas security advisories are official notifications released by the Renesas Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexRenesas(queryParameters ...IndexQueryParameters) (responseJSON *IndexRenesasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("renesas"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexReviveResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryRevive `json:"data"`
}

// GetIndexRevive -  Revive security advisories sare official notifications released by Revive to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexRevive(queryParameters ...IndexQueryParameters) (responseJSON *IndexReviveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("revive"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRocheResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryRoche `json:"data"`
}

// GetIndexRoche -  Roche security advisories are official notifications released by Roche to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexRoche(queryParameters ...IndexQueryParameters) (responseJSON *IndexRocheResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("roche"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRockwellResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryRockwell `json:"data"`
}

// GetIndexRockwell -  Rockwell Automation security advisories are official notifications released by Rockwell Automation to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexRockwell(queryParameters ...IndexQueryParameters) (responseJSON *IndexRockwellResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rockwell"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRockyResponse struct {
	Benchmark float64            `json:"_benchmark"`
	Meta      IndexMeta          `json:"_meta"`
	Data      []client.ApiUpdate `json:"data"`
}

// GetIndexRocky -  The Rocky Linux community and development team work diligently to identify and address vulnerabilities by providing regular security updates and advisories, helping to maintain a more secure environment for Rocky Linux users.

func (c *Client) GetIndexRocky(queryParameters ...IndexQueryParameters) (responseJSON *IndexRockyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rocky"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRockyErrataResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryRockyErrata `json:"data"`
}

// GetIndexRockyErrata -  Rocky Errata is a collection of official notifications released by Rocky Linux to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexRockyErrata(queryParameters ...IndexQueryParameters) (responseJSON *IndexRockyErrataResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rocky-errata"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRockyPurlsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.PurlsPurlResponse `json:"data"`
}

// GetIndexRockyPurls -  Rocky purls is a collection of rocky package purls with their associated versions and cves.

func (c *Client) GetIndexRockyPurls(queryParameters ...IndexQueryParameters) (responseJSON *IndexRockyPurlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rocky-purls"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRsyncResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryRsync `json:"data"`
}

// GetIndexRsync -  Rsync security advisories are official notifications released by the Rsync team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexRsync(queryParameters ...IndexQueryParameters) (responseJSON *IndexRsyncResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rsync"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRuckusResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryRuckus `json:"data"`
}

// GetIndexRuckus -  Ruckus security bulletins are official notifications released by Ruckus to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexRuckus(queryParameters ...IndexQueryParameters) (responseJSON *IndexRuckusResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ruckus"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexRustsecAdvisoriesResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryRustsecAdvisory `json:"data"`
}

// GetIndexRustsecAdvisories -  RustSec Advisories are security advisories filed against crates published via crates.io and are maintained by the Rust Secure Code Working Group.

func (c *Client) GetIndexRustsecAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexRustsecAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rustsec-advisories"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSacertResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySAAdvisory `json:"data"`
}

// GetIndexSacert -  Saudi CERT security alerts are official notifications released by the Saudi CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexSacert(queryParameters ...IndexQueryParameters) (responseJSON *IndexSacertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sacert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSaintResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisorySaintExploit `json:"data"`
}

// GetIndexSaint -  SAINT Exploits exploits list are advisories and contain vulnerability details that are curated by the SAINT Corporation.

func (c *Client) GetIndexSaint(queryParameters ...IndexQueryParameters) (responseJSON *IndexSaintResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("saint"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSalesforceResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySalesForce `json:"data"`
}

// GetIndexSalesforce -  SalesForce security advisories are official notifications released by SalesForce to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSalesforce(queryParameters ...IndexQueryParameters) (responseJSON *IndexSalesforceResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("salesforce"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSambaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisorySamba `json:"data"`
}

// GetIndexSamba -  Samba security releases are official notifications released by the Samba open source project to address security vulnerabilities and updates in the open source Samba project. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSamba(queryParameters ...IndexQueryParameters) (responseJSON *IndexSambaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("samba"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSapResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisorySAP `json:"data"`
}

// GetIndexSap -  SAP Security Patch Days are official notifications released by the SAP Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSap(queryParameters ...IndexQueryParameters) (responseJSON *IndexSapResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sap"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSchneiderElectricResponse struct {
	Benchmark float64                                    `json:"_benchmark"`
	Meta      IndexMeta                                  `json:"_meta"`
	Data      []client.AdvisorySchneiderElectricAdvisory `json:"data"`
}

// GetIndexSchneiderElectric -  Schneider Electric security notifications are official notifications released by Schneider Electric to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSchneiderElectric(queryParameters ...IndexQueryParameters) (responseJSON *IndexSchneiderElectricResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("schneider-electric"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSecConsultResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySECConsult `json:"data"`
}

// GetIndexSecConsult -  SEC Consult security advisories are official notifications released by SEC Consult to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSecConsult(queryParameters ...IndexQueryParameters) (responseJSON *IndexSecConsultResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sec-consult"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSecuritylabResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySecurityLab `json:"data"`
}

// GetIndexSecuritylab -  Security Lab Advisories are official notifications released by Positive Research's Security Lab to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexSecuritylab(queryParameters ...IndexQueryParameters) (responseJSON *IndexSecuritylabResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("securitylab"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSeebugResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisorySeebugExploit `json:"data"`
}

// GetIndexSeebug -  Seebug Vulnerabilities is an archive of public exploits curated by Knownsec.
func (c *Client) GetIndexSeebug(queryParameters ...IndexQueryParameters) (responseJSON *IndexSeebugResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("seebug"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSelResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisorySel `json:"data"`
}

// GetIndexSel -  Schweitzer Engineering Laboratories (SEL) security notifications are official notifications released by SEL to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSel(queryParameters ...IndexQueryParameters) (responseJSON *IndexSelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sel"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSentineloneResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySentinelOne `json:"data"`
}

// GetIndexSentinelone -  SentinelOne vulnerabilities are official notifications released by Sentinel Labs to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSentinelone(queryParameters ...IndexQueryParameters) (responseJSON *IndexSentineloneResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sentinelone"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexServicenowResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryServiceNow `json:"data"`
}

// GetIndexServicenow -  ServiceNow CVE security advisories are official notifications released by ServiceNow to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexServicenow(queryParameters ...IndexQueryParameters) (responseJSON *IndexServicenowResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("servicenow"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexShadowserverExploitedResponse struct {
	Benchmark float64                                             `json:"_benchmark"`
	Meta      IndexMeta                                           `json:"_meta"`
	Data      []client.AdvisoryShadowServerExploitedVulnerability `json:"data"`
}

// GetIndexShadowserverExploited -  Shadowserver foundation vulnerabilities contain attack statistics. Vulnerabilities are ranked according to the frequency with which exploitation attempts are made against honeypots.

func (c *Client) GetIndexShadowserverExploited(queryParameters ...IndexQueryParameters) (responseJSON *IndexShadowserverExploitedResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("shadowserver-exploited"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexShielderResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryShielder `json:"data"`
}

// GetIndexShielder -  Shielder Advisories are official notifications released by Shielder to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexShielder(queryParameters ...IndexQueryParameters) (responseJSON *IndexShielderResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("shielder"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSickResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySick `json:"data"`
}

// GetIndexSick -  SICK security advisories are official notifications released by the SICK Product Security Incident Response Team (SICK PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSick(queryParameters ...IndexQueryParameters) (responseJSON *IndexSickResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sick"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSiemensResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisorySiemensAdvisory `json:"data"`
}

// GetIndexSiemens -  Siemens security advisories are official notifications released by the Siemens ProductCERT to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSiemens(queryParameters ...IndexQueryParameters) (responseJSON *IndexSiemensResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("siemens"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSierraWirelessResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisorySierraWireless `json:"data"`
}

// GetIndexSierraWireless -  Sierra Wireless security bulletins notices are official notifications released by Sierra Wireless to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSierraWireless(queryParameters ...IndexQueryParameters) (responseJSON *IndexSierraWirelessResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sierra-wireless"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSigmahqSigmaRulesResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisorySigmaRule `json:"data"`
}

// GetIndexSigmahqSigmaRules -  Sigma Rules is a collection of rules where detection engineers, threat hunters and all defensive security practitioners collaborate on detection rules for SIEM systems.

func (c *Client) GetIndexSigmahqSigmaRules(queryParameters ...IndexQueryParameters) (responseJSON *IndexSigmahqSigmaRulesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sigmahq-sigma-rules"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSingcertResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySingCert `json:"data"`
}

// GetIndexSingcert -  CSA (Cyber Security Agency of Singapore) alerts and advisories are official notifications released by the CSA to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexSingcert(queryParameters ...IndexQueryParameters) (responseJSON *IndexSingcertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("singcert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSlackwareResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisorySlackware `json:"data"`
}

// GetIndexSlackware -  Slackware security advisories are official notifications released by the open source Slackware project to address security vulnerabilities and updates in the open source Slackware project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSlackware(queryParameters ...IndexQueryParameters) (responseJSON *IndexSlackwareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("slackware"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSolarwindsResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisorySolarWindsAdvisory `json:"data"`
}

// GetIndexSolarwinds -  SolarWinds security vulnerabilities are official notifications released by SolarWinds to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSolarwinds(queryParameters ...IndexQueryParameters) (responseJSON *IndexSolarwindsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("solarwinds"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSolrResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySolr `json:"data"`
}

// GetIndexSolr -  Solr cve reports are official notifications released by the open source Solr project to address vulnerabilities and updates in the open source Solr project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSolr(queryParameters ...IndexQueryParameters) (responseJSON *IndexSolrResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("solr"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSonatypeResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySonatype `json:"data"`
}

// GetIndexSonatype -  Sonatype security advisories are official notifications released by Sonatype to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexSonatype(queryParameters ...IndexQueryParameters) (responseJSON *IndexSonatypeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sonatype"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSonicwallResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisorySonicWallAdvisory `json:"data"`
}

// GetIndexSonicwall -  SonicWall security advisories are official notifications released by the SonicWall Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSonicwall(queryParameters ...IndexQueryParameters) (responseJSON *IndexSonicwallResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sonicwall"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSpacelabsHealthcareResponse struct {
	Benchmark float64                                      `json:"_benchmark"`
	Meta      IndexMeta                                    `json:"_meta"`
	Data      []client.AdvisorySpacelabsHealthcareAdvisory `json:"data"`
}

// GetIndexSpacelabsHealthcare -  Spacelabs security advisories are official notifications released by Spacelabs to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSpacelabsHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *IndexSpacelabsHealthcareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("spacelabs-healthcare"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSplunkResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisorySplunk `json:"data"`
}

// GetIndexSplunk -  Splunk security advisories are official notifications released by Splunk to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSplunk(queryParameters ...IndexQueryParameters) (responseJSON *IndexSplunkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("splunk"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSpringResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisorySpring `json:"data"`
}

// GetIndexSpring -  Spring security advisories are official notifications released by the VMWare Security Response team to address security vulnerabilities and updates in the open source Spring framework. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSpring(queryParameters ...IndexQueryParameters) (responseJSON *IndexSpringResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("spring"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSsdResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySSDAdvisory `json:"data"`
}

// GetIndexSsd -  SSD Secure Disclosure Advisories are official advisories released by SSD Secure Disclosure. Many advisories contain not only vulnerability details but also proof of concept code.

func (c *Client) GetIndexSsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexSsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ssd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexStormshieldResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryStormshield `json:"data"`
}

// GetIndexStormshield -  Stormshield advisories are official notifications released by Stormshield to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexStormshield(queryParameters ...IndexQueryParameters) (responseJSON *IndexStormshieldResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("stormshield"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexStrykerResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryStrykerAdvisory `json:"data"`
}

// GetIndexStryker -  Stryker security advisories are official notifications released by Stryker to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexStryker(queryParameters ...IndexQueryParameters) (responseJSON *IndexStrykerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("stryker"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSudoResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySudo `json:"data"`
}

// GetIndexSudo -  Sudo security advisories are official notifications released by the open source sudo project to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSudo(queryParameters ...IndexQueryParameters) (responseJSON *IndexSudoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sudo"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSuseResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCvrf `json:"data"`
}

// GetIndexSuse -  SUSE Security Advisories are official notifications from SUSE, a prominent open-source software company, that inform users about security vulnerabilities and provide guidance on mitigating risks in their Linux-based products and solutions. These advisories play a crucial role in helping SUSE users maintain the security and integrity of their systems.

func (c *Client) GetIndexSuse(queryParameters ...IndexQueryParameters) (responseJSON *IndexSuseResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("suse"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSuseSecurityResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisorySuseSecurity `json:"data"`
}

// GetIndexSuseSecurity -  Suse security advisories are official notifications released by the Suse Security Team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexSuseSecurity(queryParameters ...IndexQueryParameters) (responseJSON *IndexSuseSecurityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("suse-security"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSwiftResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

// GetIndexSwift -  Swift packages with package versions, associated licenses, and relevant CVEs

func (c *Client) GetIndexSwift(queryParameters ...IndexQueryParameters) (responseJSON *IndexSwiftResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("swift"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSwisslogHealthcareResponse struct {
	Benchmark float64                                     `json:"_benchmark"`
	Meta      IndexMeta                                   `json:"_meta"`
	Data      []client.AdvisorySwisslogHealthcareAdvisory `json:"data"`
}

// GetIndexSwisslogHealthcare -  Swisslog Healthcare CVE Disclosures are official notifications released by Swisslog Healthcare to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSwisslogHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *IndexSwisslogHealthcareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("swisslog-healthcare"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSymfonyResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisorySymfony `json:"data"`
}

// GetIndexSymfony -  Symfony security advisories are official notifications released by the open source Symfony project to address security vulnerabilities and updates in the open source Symfony project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSymfony(queryParameters ...IndexQueryParameters) (responseJSON *IndexSymfonyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("symfony"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSynacktivResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisorySynacktiv `json:"data"`
}

// GetIndexSynacktiv -  Synacktiv security advisories are official notifications released by Synacktiv to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSynacktiv(queryParameters ...IndexQueryParameters) (responseJSON *IndexSynacktivResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("synacktiv"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSyncrosoftResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySyncroSoft `json:"data"`
}

// GetIndexSyncrosoft -  SyncroSoft security advisories are official notifications released by SyncroSoft to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSyncrosoft(queryParameters ...IndexQueryParameters) (responseJSON *IndexSyncrosoftResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("syncrosoft"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexSynologyResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySynology `json:"data"`
}

// GetIndexSynology -  Synology product security advisories are official notifications released by Synology to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexSynology(queryParameters ...IndexQueryParameters) (responseJSON *IndexSynologyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("synology"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTailscaleResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryTailscale `json:"data"`
}

// GetIndexTailscale -  Tailscale security bulletins are official notifications released by the Tailscale team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexTailscale(queryParameters ...IndexQueryParameters) (responseJSON *IndexTailscaleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tailscale"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTeamviewerResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryTeamViewer `json:"data"`
}

// GetIndexTeamviewer -  TeamViewer security bulletins are official notifications released by TeamViewer to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTeamviewer(queryParameters ...IndexQueryParameters) (responseJSON *IndexTeamviewerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("teamviewer"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTenableResearchAdvisoriesResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryTenableResearchAdvisory `json:"data"`
}

// GetIndexTenableResearchAdvisories -  Tenable Research Advisories are official notifications released by Tenable to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexTenableResearchAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexTenableResearchAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tenable-research-advisories"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTencentResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryTencent `json:"data"`
}

// GetIndexTencent -  Tencent vulnerability risk notices are official notifications released by Tencent to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTencent(queryParameters ...IndexQueryParameters) (responseJSON *IndexTencentResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tencent"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexThalesResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryThales `json:"data"`
}

// GetIndexThales -  Thales security updates are official notifications released by Thales to address security vulnerabilities and updates in their software products. These updates provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexThales(queryParameters ...IndexQueryParameters) (responseJSON *IndexThalesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("thales"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexThemissinglinkResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryTheMissingLink `json:"data"`
}

// GetIndexThemissinglink -  the missing link security advisories are official notifications released by the missing link to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexThemissinglink(queryParameters ...IndexQueryParameters) (responseJSON *IndexThemissinglinkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("themissinglink"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexThreatActorsResponse struct {
	Benchmark float64                                         `json:"_benchmark"`
	Meta      IndexMeta                                       `json:"_meta"`
	Data      []client.AdvisoryThreatActorWithExternalObjects `json:"data"`
}

// GetIndexThreatActors -  The VulnCheck Threat Actors index contains data related to various threat actors.

func (c *Client) GetIndexThreatActors(queryParameters ...IndexQueryParameters) (responseJSON *IndexThreatActorsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("threat-actors"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTiResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryTI `json:"data"`
}

// GetIndexTi -  Texas Instrument product security bulletins are official notifications released by the Texas Instruments Product Security Incident Response Team (PSIRT) to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTi(queryParameters ...IndexQueryParameters) (responseJSON *IndexTiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ti"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTibcoResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryTibco `json:"data"`
}

// GetIndexTibco -  TIBCO security advisories are official notifications released by TIBCO to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTibco(queryParameters ...IndexQueryParameters) (responseJSON *IndexTibcoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tibco"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTpLinkResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryTPLink `json:"data"`
}

// GetIndexTpLink -  TP-Link security advisories are official notifications released by TP-Link to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTpLink(queryParameters ...IndexQueryParameters) (responseJSON *IndexTpLinkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tp-link"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTraneTechnologyResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryTraneTechnology `json:"data"`
}

// GetIndexTraneTechnology -  Trane Technology product security advisories are official notifications released by Trane Technology to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTraneTechnology(queryParameters ...IndexQueryParameters) (responseJSON *IndexTraneTechnologyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("trane-technology"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTrendmicroResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryTrendMicro `json:"data"`
}

// GetIndexTrendmicro -  Trend Micro security bulletins are official notifications released by Trend Micro to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTrendmicro(queryParameters ...IndexQueryParameters) (responseJSON *IndexTrendmicroResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("trendmicro"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTrustwaveResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryTrustwave `json:"data"`
}

// GetIndexTrustwave -  Trustwave security advisories are official notifications released by SpiderLabs to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexTrustwave(queryParameters ...IndexQueryParameters) (responseJSON *IndexTrustwaveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("trustwave"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTwcertResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryTWCertAdvisory `json:"data"`
}

// GetIndexTwcert -  Taiwan CERT vulnerability notes are official notifications released by the Taiwan CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexTwcert(queryParameters ...IndexQueryParameters) (responseJSON *IndexTwcertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("twcert"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexUbiquitiResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryUbiquiti `json:"data"`
}

// GetIndexUbiquiti -  Ubiquiti security advisorie bulletins are official notifications released by Dell to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexUbiquiti(queryParameters ...IndexQueryParameters) (responseJSON *IndexUbiquitiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ubiquiti"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexUbuntuResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryUbuntuCVE `json:"data"`
}

// GetIndexUbuntu -  Ubuntu security advisories are official notifications released by Ubuntu to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexUbuntu(queryParameters ...IndexQueryParameters) (responseJSON *IndexUbuntuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ubuntu"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexUnifyResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryUnify `json:"data"`
}

// GetIndexUnify -  Unify product security advisories and security notes are official notifications released by Unify to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexUnify(queryParameters ...IndexQueryParameters) (responseJSON *IndexUnifyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("unify"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexUnisocResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUnisoc `json:"data"`
}

// GetIndexUnisoc -  UNISOC security bulletins are official notifications released by UNISOC to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexUnisoc(queryParameters ...IndexQueryParameters) (responseJSON *IndexUnisocResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("unisoc"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexUsdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryUSD `json:"data"`
}

// GetIndexUsd -  usd advisories are official notifications released by the usd HeroLab to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexUsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexUsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("usd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexUsomResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryUSOMAdvisory `json:"data"`
}

// GetIndexUsom -  USOM security notices are official notifications released by the Turkey USOM TR-CERT to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexUsom(queryParameters ...IndexQueryParameters) (responseJSON *IndexUsomResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("usom"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVandykeResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVanDyke `json:"data"`
}

// GetIndexVandyke -  VanDyke security advisories are official notifications released by VanDyke to address security vulnerabilities and updates in VanDyke. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexVandyke(queryParameters ...IndexQueryParameters) (responseJSON *IndexVandykeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vandyke"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVapidlabsResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryVapidLabsAdvisory `json:"data"`
}

// GetIndexVapidlabs -  VapidLabs Vulnerabilities are advisories and contain vulnerability details along with exploits that are curated by Larry Cashdollar.

func (c *Client) GetIndexVapidlabs(queryParameters ...IndexQueryParameters) (responseJSON *IndexVapidlabsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vapidlabs"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVcCpeDictionaryResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryVCCPEDictionary `json:"data"`
}

// GetIndexVcCpeDictionary -  A dictionary of CPEs used in the construction of VCConfigurations.
func (c *Client) GetIndexVcCpeDictionary(queryParameters ...IndexQueryParameters) (responseJSON *IndexVcCpeDictionaryResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vc-cpe-dictionary"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVdeResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryVDEAdvisory `json:"data"`
}

// GetIndexVde -  VDE CERT Advisories are official notifications released by VDE CERT to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexVde(queryParameters ...IndexQueryParameters) (responseJSON *IndexVdeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vde"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVeeamResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryVeeam `json:"data"`
}

// GetIndexVeeam -  Veeam security advisories are official notifications released by Veeam to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexVeeam(queryParameters ...IndexQueryParameters) (responseJSON *IndexVeeamResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("veeam"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVeritasResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVeritas `json:"data"`
}

// GetIndexVeritas -  Veritas security alerts are official notifications released by Veritas to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexVeritas(queryParameters ...IndexQueryParameters) (responseJSON *IndexVeritasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("veritas"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVirtuozzoResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryVirtuozzo `json:"data"`
}

// GetIndexVirtuozzo -  Virtuozzo security advisories are official notifications released by Virtuozzo to address security vulnerabilities and updates for the Virtuozzo ReadyKernel patch service. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexVirtuozzo(queryParameters ...IndexQueryParameters) (responseJSON *IndexVirtuozzoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("virtuozzo"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVlcResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryVLC `json:"data"`
}

// GetIndexVlc -  VLC security advisories are official notifications released by VLC to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexVlc(queryParameters ...IndexQueryParameters) (responseJSON *IndexVlcResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vlc"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVmwareResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryVMWareAdvisory `json:"data"`
}

// GetIndexVmware -  VMWare security advisories are official notifications released by Broadcom to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.

func (c *Client) GetIndexVmware(queryParameters ...IndexQueryParameters) (responseJSON *IndexVmwareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vmware"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVoidsecResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVoidSec `json:"data"`
}

// GetIndexVoidsec -  VoidSec advisories are official notifications released by VoidSec to address security vulnerabilities and updates in third party products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexVoidsec(queryParameters ...IndexQueryParameters) (responseJSON *IndexVoidsecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("voidsec"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryVulnCheck `json:"data"`
}

// GetIndexVulncheck -  VulnCheck Security Advisories are official advisories released by VulnCheck to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure security.
func (c *Client) GetIndexVulncheck(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckConfigResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryVulnCheckConfig `json:"data"`
}

// GetIndexVulncheckConfig -  VulnCheck configurations contain curated/generated cpe criteria matches for a given cve based off of the Mitre CVE dataset and NVD dictionary and VulnCheck CPE dictionary.

func (c *Client) GetIndexVulncheckConfig(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckConfigResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-config"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckCvelistV5Response struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryVulnCheckCVEListV5 `json:"data"`
}

// GetIndexVulncheckCvelistV5 -  VulnCheck CVEList-V5 is a collection of publicly disclosed cybersecurity vulnerabilities by NIST that aims to identify, define and catalog publicly disclosed cybersecurity vulnerabilities. VulnCheck has curated and enhanced the data present in the NIST vulnerabilities.

func (c *Client) GetIndexVulncheckCvelistV5(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckCvelistV5Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-cvelist-v5"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckKevResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryVulnCheckKEV `json:"data"`
}

// GetIndexVulncheckKev -  The VulnCheck Known Exploit Vulnerabilities catalog contains a list of exploited vulnerabilities known to VulnCheck

func (c *Client) GetIndexVulncheckKev(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckKevResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-kev"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckNvdResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiCveItemsExtended `json:"data"`
}

// GetIndexVulncheckNvd -  NVD 2.0 CVE data formatted according to the NVD 1.0 CVE schema augmented with VulnCheck data.

func (c *Client) GetIndexVulncheckNvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckNvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-nvd"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckNvd2Response struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiNVD20CVEExtended `json:"data"`
}

// GetIndexVulncheckNvd2 -  NIST NVD CVE 2.0 API data supplemented with VulnCheck Data

func (c *Client) GetIndexVulncheckNvd2(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckNvd2Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-nvd2"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulnerabilityAliasesResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.ApiVulnerabilityAlias `json:"data"`
}

// GetIndexVulnerabilityAliases -  The Vulnerability Aliases index contains the names or aliases associated with a particular vulnerability. Examples: Log4Shell, LogJam, HeatBleed, etc.

func (c *Client) GetIndexVulnerabilityAliases(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulnerabilityAliasesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulnerability-aliases"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulnrichmentResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryVulnrichment `json:"data"`
}

// GetIndexVulnrichment -  The CISA Vulnrichment project is the public repository of CISA's enrichment of public CVE records through CISA's ADP (Authorized Data Publisher) container. In this phase of the project, CISA is assessing new and recent CVEs and adding key SSVC decision points. Once scored, some higher-risk CVEs will also receive enrichment of CWE, CVSS, and CPE data points, where possible.

func (c *Client) GetIndexVulnrichment(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulnrichmentResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulnrichment"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVyaireResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryVYAIREAdvisory `json:"data"`
}

// GetIndexVyaire -  Vyaire security bulletins are official notifications released by Vyaire to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexVyaire(queryParameters ...IndexQueryParameters) (responseJSON *IndexVyaireResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vyaire"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWatchguardResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryWatchGuard `json:"data"`
}

// GetIndexWatchguard -  WatchGuard security advisories are official notifications released by WatchGuard to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWatchguard(queryParameters ...IndexQueryParameters) (responseJSON *IndexWatchguardResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("watchguard"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWhatsappResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryWhatsApp `json:"data"`
}

// GetIndexWhatsapp -  WhatsApp security advisories are official notifications released by WhatsApp to address security vulnerabilities and updates in WhatsApp. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWhatsapp(queryParameters ...IndexQueryParameters) (responseJSON *IndexWhatsappResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("whatsapp"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWibuResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryWibu `json:"data"`
}

// GetIndexWibu -  Wibu Systems security advisories are official notifications released by Wibu Systems to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWibu(queryParameters ...IndexQueryParameters) (responseJSON *IndexWibuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wibu"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWiresharkResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryWireshark `json:"data"`
}

// GetIndexWireshark -  Wireshark security advisories are official notifications released by the open source Wireshark project to address security vulnerabilities and updates in the open source Wireshark project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWireshark(queryParameters ...IndexQueryParameters) (responseJSON *IndexWiresharkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wireshark"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWithSecureResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryWithSecure `json:"data"`
}

// GetIndexWithSecure -  With Secure security advisories are official notifications released by With Secure to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWithSecure(queryParameters ...IndexQueryParameters) (responseJSON *IndexWithSecureResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("with-secure"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWolfiResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryWolfi `json:"data"`
}

// GetIndexWolfi -  Wolfi is a new community Linux undistribution that combines the best aspects of existing container base images with default security measures that will include software signatures powered by Sigstore, provenance, and software bills of material (SBOM).

func (c *Client) GetIndexWolfi(queryParameters ...IndexQueryParameters) (responseJSON *IndexWolfiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wolfi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWolfsslResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryWolfSSL `json:"data"`
}

// GetIndexWolfssl -  WolfSSL security vulnerabilities are official notifications released by WolfSSL to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWolfssl(queryParameters ...IndexQueryParameters) (responseJSON *IndexWolfsslResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wolfssl"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexWordfenceResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryWordfence `json:"data"`
}

// GetIndexWordfence -  Wordfence vulnerabilities are official notifications released by Wordfence to address security vulnerabilities and updates in open source WordPress plugins. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexWordfence(queryParameters ...IndexQueryParameters) (responseJSON *IndexWordfenceResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wordfence"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexXenResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryXen `json:"data"`
}

// GetIndexXen -  Xen advisories are official notifications released by the open source Xen project to address vulnerabilities and updates in the open source Apache ZooKeeper project. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexXen(queryParameters ...IndexQueryParameters) (responseJSON *IndexXenResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xen"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexXeroxResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryXerox `json:"data"`
}

// GetIndexXerox -  Xerox security bulletins are official notifications released by Xerox to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexXerox(queryParameters ...IndexQueryParameters) (responseJSON *IndexXeroxResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xerox"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexXiaomiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryXiaomi `json:"data"`
}

// GetIndexXiaomi -  Xiaomi security bulletins are official notifications released by Xiaomi to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexXiaomi(queryParameters ...IndexQueryParameters) (responseJSON *IndexXiaomiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xiaomi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexXylemResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryXylem `json:"data"`
}

// GetIndexXylem -  Xylem security advisories are official notifications released by Xylem to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexXylem(queryParameters ...IndexQueryParameters) (responseJSON *IndexXylemResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xylem"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexYamahaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryYamaha `json:"data"`
}

// GetIndexYamaha -  Yamaha security advisories are official notifications released by the Yamaha team to address security vulnerabilities and updates. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.
func (c *Client) GetIndexYamaha(queryParameters ...IndexQueryParameters) (responseJSON *IndexYamahaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("yamaha"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexYokogawaResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryYokogawaAdvisory `json:"data"`
}

// GetIndexYokogawa -  Yokogawa security advisories are official notifications released by Yokogawa to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexYokogawa(queryParameters ...IndexQueryParameters) (responseJSON *IndexYokogawaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("yokogawa"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexYubicoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryYubico `json:"data"`
}

// GetIndexYubico -  Yubico security advisories are official notifications released by Yubico to address security vulnerabilities and updates in their software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexYubico(queryParameters ...IndexQueryParameters) (responseJSON *IndexYubicoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("yubico"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZdiResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryZeroDayAdvisory `json:"data"`
}

// GetIndexZdi -  Zero Day Initiative advisories are official advisories released by Trend Micro to promote responsible disclosure of vulnerabilities.

func (c *Client) GetIndexZdi(queryParameters ...IndexQueryParameters) (responseJSON *IndexZdiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zdi"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZebraResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryZebra `json:"data"`
}

// GetIndexZebra -  Zebra security alerts are official notifications released by Zebra to address security vulnerabilities and updates in their software and hardware products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexZebra(queryParameters ...IndexQueryParameters) (responseJSON *IndexZebraResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zebra"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZeroscienceResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryZeroScienceAdvisory `json:"data"`
}

// GetIndexZeroscience -  ZeroScience Vulnerabilities are vulnerability notices released by the ZeroScience Lab. Many vulnerabilities contain not only vulnerability details but also proof of concept code.

func (c *Client) GetIndexZeroscience(queryParameters ...IndexQueryParameters) (responseJSON *IndexZeroscienceResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zeroscience"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZimbraResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryZimbra `json:"data"`
}

// GetIndexZimbra -  Zimbra security advisories are official notifications released by Zimbra to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexZimbra(queryParameters ...IndexQueryParameters) (responseJSON *IndexZimbraResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zimbra"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZoomResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryZoom `json:"data"`
}

// GetIndexZoom -  Zoom security bulletins are official notifications released by Zoom to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexZoom(queryParameters ...IndexQueryParameters) (responseJSON *IndexZoomResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zoom"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZscalerResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryZscaler `json:"data"`
}

// GetIndexZscaler -  Zscaler security advisories are official notifications released by Zscaler to address security vulnerabilities and updates in third party software products. These security advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexZscaler(queryParameters ...IndexQueryParameters) (responseJSON *IndexZscalerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zscaler"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZusoResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryZuso `json:"data"`
}

// GetIndexZuso -  Zuso vulnerability notifications are official notifications released by Zuso Generation to address security vulnerabilities found in external software.

func (c *Client) GetIndexZuso(queryParameters ...IndexQueryParameters) (responseJSON *IndexZusoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zuso"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexZyxelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryZyxel `json:"data"`
}

// GetIndexZyxel -  Zyxel security advisories are official notifications released by Zyxel to address security vulnerabilities and updates in their software products. These advisories provide important information about the vulnerabilities, their potential impact, and recommendations for users to apply necessary patches or updates to ensure the security of their systems.

func (c *Client) GetIndexZyxel(queryParameters ...IndexQueryParameters) (responseJSON *IndexZyxelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zyxel"), nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		return nil, handleErrorResponse(resp)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}
