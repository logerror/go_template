package test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type MavenMetadataConfig struct {
	XMLName     xml.Name `xml:"metadata"`
	GroupId   string   `xml:"groupId"`
	ArtifactId   string   `xml:"artifactId"`
	Version     string   `xml:"version"`
	Versioning VersioningConfig `xml:"versioning"`
}

type VersioningConfig struct {
	XMLName     xml.Name `xml:"versioning"`
	LastUpdated string `xml:"lastUpdated"`
	SnapshotVersions  [] SnapshotVersionsConfig  `xml:"snapshotVersions"`
}

type SnapshotVersionsConfig struct {
	XMLName   xml.Name `xml:"snapshotVersions"`
	SnapshotVersion [] SnapshotVersionConfig `xml:"snapshotVersion"`
}

type SnapshotVersionConfig struct {
	XMLName   xml.Name `xml:"snapshotVersion"`
	Extension string `xml:"extension"`
	Value string `xml:"value"`
	Updated string `xml:"updated"`
}

func unmarshalXml(){
	reportData, _ := ioutil.ReadFile("/home/will/metadata.xml")
	reportString := string(reportData)

	metadata := MavenMetadataConfig{}
	_ = xml.Unmarshal([]byte(reportString), &metadata)
	for i, d := range metadata.Versioning.SnapshotVersions {
		fmt.Println("第", i, "个数的值是:", d.SnapshotVersion)
	}
}

//<?xml version="1.0" encoding="UTF-8"?>
//<metadata modelVersion="1.1.0">
//  <groupId>org.example</groupId>
//  <artifactId>maven_hash</artifactId>
//  <version>1.0-SNAPSHOT</version>
//  <versioning>
//    <snapshot>
//      <timestamp>20200901.070201</timestamp>
//      <buildNumber>1</buildNumber>
//    </snapshot>
//    <lastUpdated>20200901070201</lastUpdated>
//    <snapshotVersions>
//      <snapshotVersion>
//        <extension>jar</extension>
//        <value>1.0-20200901.070201-1</value>
//        <updated>20200901070201</updated>
//      </snapshotVersion>
//      <snapshotVersion>
//        <extension>pom</extension>
//        <value>1.0-20200901.070201-1</value>
//        <updated>20200901070201</updated>
//      </snapshotVersion>
//    </snapshotVersions>
//  </versioning>
//</metadata>
