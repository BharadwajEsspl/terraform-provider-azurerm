package sapavailabilityzonedetails

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SAPDatabaseType string

const (
	SAPDatabaseTypeDBTwo SAPDatabaseType = "DB2"
	SAPDatabaseTypeHANA  SAPDatabaseType = "HANA"
)

func PossibleValuesForSAPDatabaseType() []string {
	return []string{
		string(SAPDatabaseTypeDBTwo),
		string(SAPDatabaseTypeHANA),
	}
}

func (s *SAPDatabaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSAPDatabaseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSAPDatabaseType(input string) (*SAPDatabaseType, error) {
	vals := map[string]SAPDatabaseType{
		"db2":  SAPDatabaseTypeDBTwo,
		"hana": SAPDatabaseTypeHANA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SAPDatabaseType(input)
	return &out, nil
}

type SAPProductType string

const (
	SAPProductTypeECC       SAPProductType = "ECC"
	SAPProductTypeOther     SAPProductType = "Other"
	SAPProductTypeSFourHANA SAPProductType = "S4HANA"
)

func PossibleValuesForSAPProductType() []string {
	return []string{
		string(SAPProductTypeECC),
		string(SAPProductTypeOther),
		string(SAPProductTypeSFourHANA),
	}
}

func (s *SAPProductType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSAPProductType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSAPProductType(input string) (*SAPProductType, error) {
	vals := map[string]SAPProductType{
		"ecc":    SAPProductTypeECC,
		"other":  SAPProductTypeOther,
		"s4hana": SAPProductTypeSFourHANA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SAPProductType(input)
	return &out, nil
}
