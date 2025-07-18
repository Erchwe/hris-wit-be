package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/wit-id/blueprint-backend-go/common/httpservice"
	"github.com/wit-id/blueprint-backend-go/common/utility"
	"github.com/wit-id/blueprint-backend-go/src/repository/payload"
	sqlc "github.com/wit-id/blueprint-backend-go/src/repository/pgbo_sqlc"
	"github.com/wit-id/blueprint-backend-go/toolkit/log"
)

func (s *DataPelamarService) InsertDataPelamarLengkap(
	ctx context.Context,
	request payload.InsertPelamarLengkapPayload,
) (err error) {
	tx, err := s.mainDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed begin tx")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	q := sqlc.New(s.mainDB).WithTx(tx)

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				log.FromCtx(ctx).Error(err, "error rollback", rollBackErr)
			}
		}
	}()

	idPelamar := utility.GenerateGoogleUUID()

	request.Pelamar.IDDataPelamar = idPelamar

	request.Keluarga.IdPelamar = idPelamar

	for i := range request.Anak {
		request.Anak[i].IdPelamar = idPelamar
	}

	for i := range request.PendidikanFormal {
		request.PendidikanFormal[i].IdPelamar = idPelamar
	}
	for i := range request.PendidikanNonFormal {
		request.PendidikanNonFormal[i].IdPelamar = idPelamar
	}
	for i := range request.Bahasa {
		request.Bahasa[i].IdPelamar = idPelamar
	}
	for i := range request.Referensi {
		request.Referensi[i].IdPelamar = idPelamar
	}
	for i := range request.SaudaraKandung {
		request.SaudaraKandung[i].IdPelamar = idPelamar
	}

	// Insert pelamar utama
	_, err = q.CreatePelamar(ctx, request.Pelamar.ToEntity(s.cfg))
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed insert pelamar")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}
	fmt.Println("Pelamar utama berhasil disimpan")

	// Insert anak
	for _, a := range request.Anak {
		entity, err := a.ToEntity(s.cfg)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed convert entity anak")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
		_, err = q.CreatePelamarAnak(ctx, entity)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed insert anak")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}
	
	// Insert keluarga
	entityKeluarga, err := request.Keluarga.ToEntity(s.cfg)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed convert entity keluarga")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}
	_, err = q.CreatePelamarKeluarga(ctx, entityKeluarga)
	if err != nil {
		log.FromCtx(ctx).Error(err, "failed insert keluarga")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	// Insert pendidikan formal
	for _, pf := range request.PendidikanFormal {
		entity, err := pf.ToEntity(s.cfg)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed convert entity pendidikan formal")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
		_, err = q.CreatePelamarPendidikanFormal(ctx, entity)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed insert pendidikan formal")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}

	// Insert pendidikan non formal
	for _, pnf := range request.PendidikanNonFormal {
		entity, err := pnf.ToEntity(s.cfg)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed convert entity pendidikan non formal")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
		_, err = q.CreatePelamarPendidikanNonFormal(ctx, entity)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed insert pendidikan non formal")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}

	// Insert bahasa
	for _, b := range request.Bahasa {
		entity := b.ToEntity(s.cfg)
		_, err = q.CreatePelamarPenguasaanBahasa(ctx, entity)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed insert bahasa")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}

	// Insert referensi
	for _, r := range request.Referensi {
		entity := r.ToEntity(s.cfg)
		_, err = q.CreatePelamarReferensi(ctx, entity)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed insert referensi")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}

	// Insert saudara kandung
	for _, sdr := range request.SaudaraKandung {
		entity, err := sdr.ToEntity(s.cfg)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed convert entity saudara kandung")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
		_, err = q.CreatePelamarSaudaraKandung(ctx, entity)
		if err != nil {
			log.FromCtx(ctx).Error(err, "failed insert saudara kandung")
			return errors.WithStack(httpservice.ErrUnknownSource)
		}
	}


	if err = tx.Commit(); err != nil {
		log.FromCtx(ctx).Error(err, "error commit")
		return errors.WithStack(httpservice.ErrUnknownSource)
	}

	fmt.Println("Transaksi pelamar lengkap berhasil di-commit")
	return nil
}
