package service

//// Function to get the last lowongan pekerjaan ID
//func (s *LowonganPekerjaanService) GetLastLowonganPekerjaanID(ctx context.Context) (string, error) {
//	q := sqlc.New(s.mainDB)
//
//	var lastID string
//	err, _ := q.GetLastLowonganPekerjaanID(ctx, string(lastID))
//	if err != nil {
//		if err == "id tidak ditemukan" {
//			return "", nil // Tidak ditemukan, ID pertama dimulai dari "LOW-001"
//		}
//	}
//
//	return lastID, nil
//}
//
//// Function to generate the lowongan pekerjaan ID with the format "LOW-001"
//func (s *LowonganPekerjaanService) generateLowonganID(ctx context.Context) (string, error) {
//	q := sqlc.New(s.mainDB)
//
//	// Get the latest ID from the table
//	var lastID string
//	err, _ := q.GetLastLowonganPekerjaanID(ctx, string(lastID))
//	if err != nil && err != "id tidak ditemukan" {
//		log.FromCtx(ctx).Error(err, "failed to get last lowongan pekerjaan ID")
//	}
//
//	if lastID == "" {
//		return "LOW-001", nil
//	}
//
//	lastNumber, err := strconv.Atoi(lastID[4:])
//	if err != nil {
//		log.FromCtx(ctx).Error(err, "failed to parse last ID number")
//		return "", err
//	}
//
//	newNumber := lastNumber + 1
//	return fmt.Sprintf("LOW-%03d", newNumber), nil
//}
