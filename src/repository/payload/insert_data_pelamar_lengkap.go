package payload

import "fmt"

type InsertPelamarLengkapPayload struct {
	Pelamar             InsertDataPelamarPayload                  `json:"pelamar"`
	Keluarga            InsertPelamarKeluargaPayload              `json:"keluarga"`
	Anak                []InsertPelamarAnakPayload                `json:"anak"`
	PendidikanFormal    []InsertPelamarPendidikanFormalPayload    `json:"pendidikan_formal"`
	PendidikanNonFormal []InsertPelamarPendidikanNonFormalPayload `json:"pendidikan_non_formal"`
	Bahasa              []InsertPelamarPenguasaanBahasaPayload    `json:"bahasa"`
	Referensi           []InsertPelamarReferensiPayload           `json:"referensi"`
	SaudaraKandung      []InsertPelamarSaudaraKandungPayload      `json:"saudara_kandung"`
}

func (p *InsertPelamarLengkapPayload) Validate() error {
	// Validasi pelamar utama
	if err := p.Pelamar.Validate(); err != nil {
		return err
	}
	fmt.Println("a")

	// Validasi keluarga
	if err := p.Keluarga.Validate(); err != nil {
		return err
	}

	// Validasi anak
	for _, a := range p.Anak {
		if err := a.Validate(); err != nil {
			return err
		}
	}

	// Validasi pendidikan formal
	for _, pf := range p.PendidikanFormal {
		if err := pf.Validate(); err != nil {
			return err
		}
	}

	// Validasi pendidikan non-formal
	for _, pnf := range p.PendidikanNonFormal {
		if err := pnf.Validate(); err != nil {
			return err
		}
	}

	// Validasi bahasa
	for _, b := range p.Bahasa {
		if err := b.Validate(); err != nil {
			return err
		}
	}

	// Validasi referensi
	for _, r := range p.Referensi {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	// Validasi saudara kandung
	for _, sdr := range p.SaudaraKandung {
		if err := sdr.Validate(); err != nil {
			return err
		}
	}

	return nil
}
