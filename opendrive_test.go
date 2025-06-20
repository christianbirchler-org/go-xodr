package xodr

import (
	"encoding/xml"
	"testing"
)

func TestNumberOfStructs(t *testing.T) {
	t.Skip("test not implemented")
}

func TestParseAndSetializeShouldBeEqual(t *testing.T) {
	// Metamorphic testing

	var od OpenDrive
	err := xml.Unmarshal([]byte(sampleXodr), &od)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(od.Header.RevMinor)
	t.Log(od.Road[0].Name)

	data, err := xml.Marshal(&od)
	if err != nil {
		t.Fatal(err)
	}

	var odCompare OpenDrive
	err = xml.Unmarshal(data, &odCompare)
	if err != nil {
		t.Fatal(err)
	}

	// we need this to mitigate the differences of the xml formatting
	dataCompare, err := xml.Marshal(&odCompare)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != string(dataCompare) {
		t.Fatal("should be equal", "data", data, "dataCompare", dataCompare)
	}

}

var sampleXodr = `
<?xml version="1.0" standalone="yes"?>
<OpenDRIVE>
    <header revMajor="1" revMinor="8" name="" version="1.00" date="Wed Jan 11 16:56:17 2023" north="0.0000000000000000e+00" south="0.0000000000000000e+00" east="0.0000000000000000e+00" west="0.0000000000000000e+00">
    </header>
    <road name="drivingRoad" length="1.2000000000000000e+02" id="1" junction="-1" rule="RHT">
        <link>
        </link>
        <planView>
            <geometry s="0.0000000000000000e+00" x="7.0000000000000000e+01" y="3.6000000000000000e+02" hdg="0.0000000000000000e+00" length="1.2000000000000000e+02">
                <line/>
            </geometry>
        </planView>
        <elevationProfile>
            <elevation s="0.0000000000000000e+00" a="0.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
        </elevationProfile>
        <lateralProfile>
        </lateralProfile>
        <lanes>
            <laneSection s="0.0000000000000000e+00">
                <left>
                    <lane id="4" type="walking" level="false">
                        <link>
                            <successor id="5"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="2.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="3" type="curb" level="false">
                        <link>
                            <successor id="4"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="1.4999999999999999e-01" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="2" type="driving" level="false">
                        <link>
                            <successor id="3"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                    <lane id="1" type="driving" level="false">
                        <link>
                            <successor id="2"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                </left>
                <center>
                    <lane id="0">
                        <link>
                        </link>
                    </lane>
                </center>
                <right>
                    <lane id="-1" type="driving" level="false">
                        <link>
                            <successor id="-1"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                    <lane id="-2" type="curb" level="false">
                        <link>
                            <successor id="-2"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="1.4999999999999999e-01" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="-3" type="walking" level="false">
                        <link>
                            <successor id="-3"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="2.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                </right>
            </laneSection>
            <laneSection s="4.0000000000000000e+01">
                <left>
                    <lane id="5" type="walking" level="false">
                        <link>
                            <predecessor id="4"/>
                            <successor id="4"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="2.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.7000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.8000000000000000e+01" inner="0.0000000000000000e+00" outer="1.2000000000000000e-01"/>
                        <height sOffset="2.1000000000000000e+01" inner="0.0000000000000000e+00" outer="1.2000000000000000e-01"/>
                        <height sOffset="2.2000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="4" type="curb" level="false">
                        <link>
                            <predecessor id="3"/>
                            <successor id="3"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="1.4999999999999999e-01" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.7000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.8000000000000000e+01" inner="0.0000000000000000e+00" outer="0.0000000000000000e+00"/>
                        <height sOffset="2.1000000000000000e+01" inner="0.0000000000000000e+00" outer="0.0000000000000000e+00"/>
                        <height sOffset="2.2000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="3" type="driving" level="false">
                        <link>
                            <predecessor id="2"/>
                            <successor id="2"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                    <lane id="2" type="driving" level="false">
                        <link>
                            <predecessor id="1"/>
                            <successor id="1"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="-2.7749999999999997e-02" d="9.2499999999999993e-04"/>
                        <width sOffset="2.0000000000000000e+01" a="0.0000000000000000e+00" b="0.0000000000000000e+00" c="2.7749999999999997e-02" d="-9.2499999999999993e-04"/>
                    </lane>
                    <lane id="1" type="restricted" level="false">
                        <link>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="0.0000000000000000e+00" b="0.0000000000000000e+00" c="2.7749999999999997e-02" d="-9.2499999999999993e-04"/>
                        <width sOffset="2.0000000000000000e+01" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="-2.7749999999999997e-02" d="9.2499999999999993e-04"/>
                    </lane>
                </left>
                <center>
                    <lane id="0">
                        <link>
                        </link>
                    </lane>
                </center>
                <right>
                    <lane id="-1" type="driving" level="false">
                        <link>
                            <predecessor id="-1"/>
                            <successor id="-1"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                    <lane id="-2" type="curb" level="false">
                        <link>
                            <predecessor id="-2"/>
                            <successor id="-2"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="1.4999999999999999e-01" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.7000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.8000000000000000e+01" inner="0.0000000000000000e+00" outer="0.0000000000000000e+00"/>
                        <height sOffset="2.1000000000000000e+01" inner="0.0000000000000000e+00" outer="0.0000000000000000e+00"/>
                        <height sOffset="2.2000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="-3" type="walking" level="false">
                        <link>
                            <predecessor id="-3"/>
                            <successor id="-3"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="2.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.7000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                        <height sOffset="1.8000000000000000e+01" inner="0.0000000000000000e+00" outer="1.2000000000000000e-01"/>
                        <height sOffset="2.1000000000000000e+01" inner="0.0000000000000000e+00" outer="1.2000000000000000e-01"/>
                        <height sOffset="2.2000000000000000e+01" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                </right>
            </laneSection>
            <laneSection s="8.0000000000000000e+01">
                <left>
                    <lane id="4" type="walking" level="false">
                        <link>
                            <predecessor id="5"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="2.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="3" type="curb" level="false">
                        <link>
                            <predecessor id="4"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="1.4999999999999999e-01" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="2" type="driving" level="false">
                        <link>
                            <predecessor id="3"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                    <lane id="1" type="driving" level="false">
                        <link>
                            <predecessor id="2"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                </left>
                <center>
                    <lane id="0">
                        <link>
                        </link>
                    </lane>
                </center>
                <right>
                    <lane id="-1" type="driving" level="false">
                        <link>
                            <predecessor id="-1"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.7000000000000002e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                    <lane id="-2" type="curb" level="false">
                        <link>
                            <predecessor id="-2"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="1.4999999999999999e-01" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                    <lane id="-3" type="walking" level="false">
                        <link>
                            <predecessor id="-3"/>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="2.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                        <height sOffset="0.0000000000000000e+00" inner="1.2000000000000000e-01" outer="1.2000000000000000e-01"/>
                    </lane>
                </right>
            </laneSection>
        </lanes>
        <objects>
            <object type="trafficIsland" subtype="island" name="ExampleIsland" id="0" s="5.0000000000000441e+01" t="4.4053649617126212e-13" zOffset="0.0000000000000000e+00" orientation="none" length="5.0000000000000000e+00" width="1.0000000000000000e+00" height="1.0000000000000000e-01" hdg="0.0000000000000000e+00" pitch="0.0000000000000000e+00" roll="0.0000000000000000e+00">
                <outlines>
                    <outline id="50" fillType="cobble" closed="true">
                        <cornerRoad s="52.5" t="1.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.6" t="1.1" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.7" t="0.7" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.8" t="0.6" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.9" t="0.55" dz="0.0" height="0.1"/>
                        <cornerRoad s="53.0" t="0.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="57.0" t="0.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="57.5" t="0.5" dz="0.0" height="0.01"/>
                        <cornerRoad s="61.5" t="0.5" dz="0.0" height="0.01"/>
                        <cornerRoad s="62.0" t="0.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.0" t="0.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.1" t="0.55" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.2" t="0.6" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.3" t="0.7" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.4" t="1.1" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.5" t="1.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.4" t="1.9" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.3" t="2.3" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.2" t="2.4" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.1" t="2.45" dz="0.0" height="0.1"/>
                        <cornerRoad s="66.0" t="2.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="62.0" t="2.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="61.5" t="2.5" dz="0.0" height="0.01"/>
                        <cornerRoad s="57.5" t="2.5" dz="0.0" height="0.01"/>
                        <cornerRoad s="57.0" t="2.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="53.0" t="2.5" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.9" t="2.45" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.8" t="2.4" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.7" t="2.3" dz="0.0" height="0.1"/>
                        <cornerRoad s="52.6" t="1.9" dz="0.0" height="0.1"/>
                    </outline>
                </outlines>
                <borders>
                    <border width="0.1" type="curb" outlineId="50" useCompleteOutline="true"/>
                </borders>
            </object>
        </objects>
        <signals>
        </signals>
        <surface>
        </surface>
    </road>
    <road name="pedestrian" length="1.4000000000000000e+01" id="2" junction="555" rule="RHT">
        <link>
        </link>
        <planView>
            <geometry s="0.0000000000000000e+00" x="1.2797566674362665e+02" y="3.5495121666281869e+02" hdg="1.5707963267948966e+00" length="1.4000000000000000e+01">
                <line/>
            </geometry>
        </planView>
        <elevationProfile>
            <elevation s="0.0000000000000000e+00" a="0.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
        </elevationProfile>
        <lateralProfile>
        </lateralProfile>
        <lanes>
            <laneSection s="0.0000000000000000e+00">
                <center>
                    <lane id="0">
                        <link>
                        </link>
                    </lane>
                </center>
                <right>
                    <lane id="-1" type="walking" level="false">
                        <link>
                        </link>
                        <width sOffset="0.0000000000000000e+00" a="3.0000000000000000e+00" b="0.0000000000000000e+00" c="0.0000000000000000e+00" d="0.0000000000000000e+00"/>
                    </lane>
                </right>
            </laneSection>
        </lanes>
        <objects>
        </objects>
        <signals>
        </signals>
        <surface>
        </surface>
    </road>
    <junction name="pedestrianCrossPath" type="virtual" id="555" mainRoad="1" sStart="57.5" sEnd="61.5">
        <crossPath id="0" crossingRoad="2" roadAtStart="1" roadAtEnd="1">
            <startLaneLink s="5.9000000000000000e+01" from="4" to="-1"/>
            <endLaneLink s="5.90000000000000000e+01" from="-3" to="-1"/>
        </crossPath>
        <priority high="1" low="2"/>
    </junction>
</OpenDRIVE>
`
