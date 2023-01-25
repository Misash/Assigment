package main

import (
	"fmt"
)

func GetUniqueFeeRecipients(FeeRecipients []string) []string {

	key := make(map[string]bool)
	uniqueFeeRecipients := []string{}

	for _, fee_recipient := range FeeRecipients {

		if !key[fee_recipient] {
			key[fee_recipient] = true
			uniqueFeeRecipients = append(uniqueFeeRecipients, fee_recipient)
		}
	}

	return uniqueFeeRecipients
}

func main() {

	data := []string{
		"0x455e5aa18469bc6ccef49594645666c587a3571b",
		"0x0069c9017BDd6753467c138449eF98320be1a4E4",
		"0x005CD1608e40d1e775a97d12e4f594029567C071",
		"0x455e5aa18469bc6ccef49594645666c587a3571b",
		"0xa891ac2337b801a95d3a0910bdeacda5c7b89234",
		"0x455e5aa18469bc6ccef49594645666c587a3571b",
		"0x8b7f0977bb4f0fbe7076fa22bc24aca043583f5e",
		"0x005CD1608e40d1e775a97d12e4f594029567C071",
		"0x005CD1608e40d1e775a97d12e4f594029567C071",
		"0x455e5aa18469bc6ccef49594645666c587a3571b",
		"0x004b1EaBc3ea60331a01fFfC3D63E5F6B3aB88B3",
		"0x455e5aa18469bc6ccef49594645666c587a3571b",
		"0x005CD1608e40d1e775a97d12e4f594029567C071",
		"0x455e5aa18469bc6ccef49594645666c587a3571b",
		"0x005CD1608e40d1e775a97d12e4f594029567C071",
	}

	FeeRecipients := GetUniqueFeeRecipients(data)

	for _, fee_recipient := range FeeRecipients {
		fmt.Println(fee_recipient)
	}

}
