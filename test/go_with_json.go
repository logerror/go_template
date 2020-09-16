package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CommonVulnerability struct {
	VulnerabilityId  string `json:"vulnerabilityId"`
	Severity         string `json:"severity"`
	PkgName          string `json:"pkgName"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	InstalledVersion string `json:"installedVersion"`
	FixedVersion     string `json:"fixedVersion"`
}

type DcVulnerabilities struct {
	VulnerabilityId  string `json:"name"`
	Severity         string `json:"severity"`
	Description      string `json:"description"`
}

type DcDependencies struct {
	Vulnerabilities [] *DcVulnerabilities `json:"vulnerabilities"`
}

type DependencyCheckReportData struct {
	//ReportSchema string `json:"reportSchema"`
	Dependencies [] DcDependencies `json:"dependencies"`
}

func unmarshalJson (){
	reportData, _ := ioutil.ReadFile("/home/will/dependency-check.json")
	reportString := string(reportData)
	var packageVulns DependencyCheckReportData

	_ = json.Unmarshal([]byte(reportString), &packageVulns)
	var vulnerabilities []*CommonVulnerability
		for _, d := range packageVulns.Dependencies {
			for _, v := range d.Vulnerabilities {
				vulnerabilities = append(vulnerabilities, &CommonVulnerability{
					VulnerabilityId: v.VulnerabilityId,
					Severity:        v.Severity,
					Description:     v.Description,
				})
			}

		}

	for i,p:=range vulnerabilities {
			fmt.Println("第", i, "个数的值是:", p.VulnerabilityId)
	}
}

//{
//"reportSchema": "1.1",
//"scanInfo": {
//"engineVersion": "5.3.2",
//"dataSource": [
//{
//"name": "NVD CVE Checked",
//"timestamp": "2020-08-10T17:35:39"
//},
//{
//"name": "NVD CVE Modified",
//"timestamp": "2020-08-10T13:03:07"
//},
//{
//"name": "VersionCheckOn",
//"timestamp": "2020-07-27T16:44:49"
//}
//]
//},
//"projectInfo": {
//"name": "",
//"reportDate": "2020-08-10T09:47:22.558Z",
//"credits": {
//"NVD": "This report contains data retrieved from the National Vulnerability Database: http://nvd.nist.gov",
//"NPM": "This report may contain data retrieved from the NPM Public Advisories: https://www.npmjs.com/advisories",
//"RETIREJS": "This report may contain data retrieved from the RetireJS community: https://retirejs.github.io/retire.js/",
//"OSSINDEX": "This report may contain data retrieved from the Sonatype OSS Index: https://ossindex.sonatype.org"
//}
//},
//"dependencies": [
//{
//"isVirtual": false,
//"fileName": "xstream-1.4.7-jenkins-1.jar",
//"filePath": "/home/will/.m2/repository/org/jvnet/hudson/xstream/1.4.7-jenkins-1/xstream-1.4.7-jenkins-1.jar",
//"md5": "6b27008bd6cb5f4cc430e219d785313a",
//"sha1": "161ed1603117c2d37b864f81a0d62f36cf7e958a",
//"sha256": "405fdd4c2e594756d2e7948acbef3b1cbe13fb024dc441a2fc8d492deb48cec3",
//"description": "XStream is a serialization library from Java objects to XML and back.",
//"license": "http://xstream.codehaus.org/license.html",
//"evidenceCollected": {
//"vendorEvidence": [
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-build-time",
//"value": "20140321-1150"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-builder",
//"value": "Maven 3.0.4"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "pom",
//"name": "artifactid",
//"value": "xstream"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-compile-target",
//"value": "1.5"
//},
//{
//"type": "vendor",
//"confidence": "HIGH",
//"source": "file",
//"name": "name",
//"value": "xstream"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "specification-vendor",
//"value": "XStream"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "pom",
//"name": "parent-artifactid",
//"value": "xstream-parent"
//},
//{
//"type": "vendor",
//"confidence": "HIGHEST",
//"source": "pom",
//"name": "groupid",
//"value": "jvnet.hudson"
//},
//{
//"type": "vendor",
//"confidence": "HIGH",
//"source": "pom",
//"name": "name",
//"value": "XStream Core"
//},
//{
//"type": "vendor",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "xstream"
//},
//{
//"type": "vendor",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "thoughtworks"
//},
//{
//"type": "vendor",
//"confidence": "MEDIUM",
//"source": "Manifest",
//"name": "bundle-symbolicname",
//"value": "xstream"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-build-os",
//"value": "Linux"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "bundle-docurl",
//"value": "http://xstream.codehaus.org"
//},
//{
//"type": "vendor",
//"confidence": "HIGH",
//"source": "Manifest",
//"name": "Implementation-Vendor",
//"value": "XStream"
//},
//{
//"type": "vendor",
//"confidence": "MEDIUM",
//"source": "Manifest",
//"name": "Implementation-Vendor-Id",
//"value": "org.jvnet.hudson"
//},
//{
//"type": "vendor",
//"confidence": "MEDIUM",
//"source": "pom",
//"name": "parent-groupid",
//"value": "com.thoughtworks.xstream"
//},
//{
//"type": "vendor",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "core"
//},
//{
//"type": "vendor",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-compile-source",
//"value": "1.5"
//}
//],
//"productEvidence": [
//{
//"type": "product",
//"confidence": "HIGH",
//"source": "Manifest",
//"name": "Implementation-Title",
//"value": "XStream Core"
//},
//{
//"type": "product",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-build-time",
//"value": "20140321-1150"
//},
//{
//"type": "product",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-builder",
//"value": "Maven 3.0.4"
//},
//{
//"type": "product",
//"confidence": "MEDIUM",
//"source": "Manifest",
//"name": "specification-title",
//"value": "XStream Core"
//},
//{
//"type": "product",
//"confidence": "MEDIUM",
//"source": "Manifest",
//"name": "Bundle-Name",
//"value": "XStream Core"
//},
//{
//"type": "product",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-compile-target",
//"value": "1.5"
//},
//{
//"type": "product",
//"confidence": "HIGH",
//"source": "file",
//"name": "name",
//"value": "xstream"
//},
//{
//"type": "product",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "xml"
//},
//{
//"type": "product",
//"confidence": "HIGHEST",
//"source": "pom",
//"name": "groupid",
//"value": "jvnet.hudson"
//},
//{
//"type": "product",
//"confidence": "HIGHEST",
//"source": "pom",
//"name": "artifactid",
//"value": "xstream"
//},
//{
//"type": "product",
//"confidence": "HIGH",
//"source": "pom",
//"name": "name",
//"value": "XStream Core"
//},
//{
//"type": "product",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "xstream"
//},
//{
//"type": "product",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "thoughtworks"
//},
//{
//"type": "product",
//"confidence": "MEDIUM",
//"source": "Manifest",
//"name": "bundle-symbolicname",
//"value": "xstream"
//},
//{
//"type": "product",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-build-os",
//"value": "Linux"
//},
//{
//"type": "product",
//"confidence": "MEDIUM",
//"source": "pom",
//"name": "parent-artifactid",
//"value": "xstream-parent"
//},
//{
//"type": "product",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "bundle-docurl",
//"value": "http://xstream.codehaus.org"
//},
//{
//"type": "product",
//"confidence": "MEDIUM",
//"source": "pom",
//"name": "parent-groupid",
//"value": "com.thoughtworks.xstream"
//},
//{
//"type": "product",
//"confidence": "HIGHEST",
//"source": "jar",
//"name": "package name",
//"value": "core"
//},
//{
//"type": "product",
//"confidence": "LOW",
//"source": "Manifest",
//"name": "x-compile-source",
//"value": "1.5"
//}
//],
//"versionEvidence": [
//{
//"type": "version",
//"confidence": "HIGHEST",
//"source": "pom",
//"name": "version",
//"value": "1.4.7-jenkins-1"
//},
//{
//"type": "version",
//"confidence": "LOW",
//"source": "pom",
//"name": "parent-version",
//"value": "1.4.7-jenkins-1"
//},
//{
//"type": "version",
//"confidence": "HIGH",
//"source": "Manifest",
//"name": "Implementation-Version",
//"value": "1.4.7-jenkins-1"
//}
//]
//},
//"packages": [
//{
//"id": "pkg:maven/org.jvnet.hudson/xstream@1.4.7-jenkins-1",
//"confidence": "HIGH",
//"url": "https://ossindex.sonatype.org/component/pkg:maven/org.jvnet.hudson/xstream@1.4.7-jenkins-1"
//}
//],
//"vulnerabilityIds": [
//{
//"id": "cpe:2.3:a:xstream_project:xstream:1.4.7.ins-1:*:*:*:*:*:*:*",
//"confidence": "HIGHEST",
//"url": "https://nvd.nist.gov/vuln/search/results?form_type=Advanced&results_type=overview&search_type=all&cpe_vendor=cpe%3A%2F%3Axstream_project&cpe_product=cpe%3A%2F%3Axstream_project%3Axstream&cpe_version=cpe%3A%2F%3Axstream_project%3Axstream%3A1.4.7.ins-1"
//}
//],
//"vulnerabilities": [
//{
//"source": "NVD",
//"name": "CVE-2016-3674",
//"severity": "HIGH",
//"cvssv2": {
//"score": 5.0,
//"accessVector": "NETWORK",
//"accessComplexity": "LOW",
//"authenticationr": "NONE",
//"confidentialImpact": "PARTIAL",
//"integrityImpact": "PARTIAL",
//"availabilityImpact": "NONE",
//"severity": "MEDIUM"
//},
//"cvssv3": {
//"baseScore": 7.5,
//"attackVector": "NETWORK",
//"attackComplexity": "LOW",
//"privilegesRequired": "NONE",
//"userInteraction": "NONE",
//"scope": "UNCHANGED",
//"confidentialityImpact": "HIGH",
//"integrityImpact": "NONE",
//"availabilityImpact": "NONE",
//"baseSeverity": "HIGH"
//},
//"cwes": [
//"CWE-200"
//],
//"description": "Multiple XML external entity (XXE) vulnerabilities in the (1) Dom4JDriver, (2) DomDriver, (3) JDomDriver, (4) JDom2Driver, (5) SjsxpDriver, (6) StandardStaxDriver, and (7) WstxDriver drivers in XStream before 1.4.9 allow remote attackers to read arbitrary files via a crafted XML document.",
//"notes": "",
//"references": [
//{
//"source": "REDHAT",
//"url": "http://rhn.redhat.com/errata/RHSA-2016-2823.html",
//"name": "RHSA-2016:2823"
//},
//{
//"source": "MLIST",
//"url": "http://www.openwall.com/lists/oss-security/2016/03/25/8",
//"name": "[oss-security] 20160325 CVE request - XStream: XXE vulnerability"
//},
//{
//"source": "FEDORA",
//"url": "http://lists.fedoraproject.org/pipermail/package-announce/2016-April/183208.html",
//"name": "FEDORA-2016-250042b8a6"
//},
//{
//"source": "SECTRACK",
//"url": "http://www.securitytracker.com/id/1036419",
//"name": "1036419"
//},
//{
//"source": "CONFIRM",
//"url": "https://github.com/x-stream/xstream/issues/25",
//"name": "https://github.com/x-stream/xstream/issues/25"
//},
//{
//"source": "MLIST",
//"url": "http://www.openwall.com/lists/oss-security/2016/03/28/1",
//"name": "[oss-security] 20160328 Re: CVE request - XStream: XXE vulnerability"
//},
//{
//"source": "BID",
//"url": "http://www.securityfocus.com/bid/85381",
//"name": "85381"
//},
//{
//"source": "CONFIRM",
//"url": "http://x-stream.github.io/changes.html#1.4.9",
//"name": "http://x-stream.github.io/changes.html#1.4.9"
//},
//{
//"source": "FEDORA",
//"url": "http://lists.fedoraproject.org/pipermail/package-announce/2016-April/183180.html",
//"name": "FEDORA-2016-de909cc333"
//},
//{
//"source": "REDHAT",
//"url": "http://rhn.redhat.com/errata/RHSA-2016-2822.html",
//"name": "RHSA-2016:2822"
//},
//{
//"source": "DEBIAN",
//"url": "http://www.debian.org/security/2016/dsa-3575",
//"name": "DSA-3575"
//}
//],
//"vulnerableSoftware": [
//{
//"software": {
//"id": "cpe:2.3:a:xstream_project:xstream:*:*:*:*:*:*:*:*",
//"vulnerabilityIdMatched": "true",
//"versionEndIncluding": "1.4.8"
//}
//}
//]
//},
//{
//"source": "NVD",
//"name": "CVE-2017-7957",
//"severity": "HIGH",
//"cvssv2": {
//"score": 5.0,
//"accessVector": "NETWORK",
//"accessComplexity": "LOW",
//"authenticationr": "NONE",
//"confidentialImpact": "NONE",
//"integrityImpact": "NONE",
//"availabilityImpact": "PARTIAL",
//"severity": "MEDIUM"
//},
//"cvssv3": {
//"baseScore": 7.5,
//"attackVector": "NETWORK",
//"attackComplexity": "LOW",
//"privilegesRequired": "NONE",
//"userInteraction": "NONE",
//"scope": "UNCHANGED",
//"confidentialityImpact": "NONE",
//"integrityImpact": "NONE",
//"availabilityImpact": "HIGH",
//"baseSeverity": "HIGH"
//},
//"cwes": [
//"CWE-20"
//],
//"description": "XStream through 1.4.9, when a certain denyTypes workaround is not used, mishandles attempts to create an instance of the primitive type 'void' during unmarshalling, leading to a remote application crash, as demonstrated by an xstream.fromXML(\"<void/>\") call.",
//"notes": "",
//"references": [
//{
//"source": "REDHAT",
//"url": "https://access.redhat.com/errata/RHSA-2017:2888",
//"name": "RHSA-2017:2888"
//},
//{
//"source": "REDHAT",
//"url": "https://access.redhat.com/errata/RHSA-2017:1832",
//"name": "RHSA-2017:1832"
//},
//{
//"source": "CONFIRM",
//"url": "https://www-prd-trops.events.ibm.com/node/715749",
//"name": "https://www-prd-trops.events.ibm.com/node/715749"
//},
//{
//"source": "SECTRACK",
//"url": "http://www.securitytracker.com/id/1039499",
//"name": "1039499"
//},
//{
//"source": "DEBIAN",
//"url": "http://www.debian.org/security/2017/dsa-3841",
//"name": "DSA-3841"
//},
//{
//"source": "BID",
//"url": "http://www.securityfocus.com/bid/100687",
//"name": "100687"
//},
//{
//"source": "CONFIRM",
//"url": "http://x-stream.github.io/CVE-2017-7957.html",
//"name": "http://x-stream.github.io/CVE-2017-7957.html"
//},
//{
//"source": "REDHAT",
//"url": "https://access.redhat.com/errata/RHSA-2017:2889",
//"name": "RHSA-2017:2889"
//},
//{
//"source": "XF",
//"url": "https://exchange.xforce.ibmcloud.com/vulnerabilities/125800",
//"name": "xstream-cve20177957-dos(125800)"
//}
//],
//"vulnerableSoftware": [
//{
//"software": {
//"id": "cpe:2.3:a:xstream_project:xstream:*:*:*:*:*:*:*:*",
//"vulnerabilityIdMatched": "true",
//"versionEndIncluding": "1.4.9"
//}
//}
//]
//}
//]
//}
//]
//}