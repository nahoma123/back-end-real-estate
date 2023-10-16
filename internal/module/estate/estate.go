package estate

import (
	"context"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/errors"
	"visitor_management/internal/constant/model"
	"visitor_management/internal/module"
	"visitor_management/internal/storage"
	"visitor_management/platform/logger"

	"go.uber.org/zap"
)

type estate struct {
	logger          logger.Logger
	estateStorage   storage.EstateStorage
	residentStorage storage.ResidentStorage
}

func InitEstate(logger logger.Logger, estateStorage storage.EstateStorage, residentStorage storage.ResidentStorage) module.EstateModule {
	return &estate{
		logger,
		estateStorage,
		residentStorage,
	}
}

func (es *estate) CreateEstate(ctx context.Context, estate *model.Estate) (*model.Estate, error) {
	if err := estate.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	estate, err := es.estateStorage.CreateEstate(ctx, estate)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return estate, nil
}

func (es *estate) CreateFacultyService(ctx context.Context, fService *model.FacultyService) (*model.FacultyService, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.estateStorage.CreateFacultyService(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) AddEstateStaff(ctx context.Context, fService *model.Staff) (*model.Staff, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.estateStorage.AddEstateStaff(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) AddHouseOwner(ctx context.Context, fService *model.HouseOwner) (*model.HouseOwner, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.estateStorage.AddHouseOwner(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) AddHouse(ctx context.Context, fService *model.House) (*model.House, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.estateStorage.AddHouse(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) GetHouses(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.House, error) {
	houses, err := es.estateStorage.GetHouses(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return houses, nil
}

func (es *estate) GetFacultyServices(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.FacultyService, error) {
	fsc, err := es.estateStorage.GetFacultyServices(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) GetEstateStaffs(ctx context.Context, filterPagination *constant.FilterPagination) ([]interface{}, error) {
	fsc, err := es.estateStorage.GetEstateStaffs(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) GetResidents(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Resident, error) {
	fsc, err := es.residentStorage.GetResidents(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) AssignResident(ctx context.Context, fService *model.Resident) (*model.Resident, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.residentStorage.AssignResident(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) GetHouseFees(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.HouseFee, error) {
	fsc, err := es.residentStorage.GetHouseFees(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) CreateHouseFee(ctx context.Context, fService *model.HouseFee) (*model.HouseFee, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.residentStorage.CreateHouseFee(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil

}
func (es *estate) GetHouseFeeSubscription(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.HouseFeeSubscription, error) {
	fsc, err := es.residentStorage.GetHouseFeeSubscription(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) AddHouseFeeSubscription(ctx context.Context, fService *model.HouseFeeSubscription) (*model.HouseFeeSubscription, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}
	fService.InitializePaymentDates()

	fService, err := es.residentStorage.AddHouseFeeSubscription(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil

}

func (es *estate) GetFacultyServiceSubscriptions(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.FacultyServiceSubscription, error) {
	fsc, err := es.estateStorage.GetFacultyServiceSubscriptions(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) AddFacultyServiceSubscription(ctx context.Context, fService *model.FacultyServiceSubscription) (*model.FacultyServiceSubscription, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}
	fService.InitializePaymentDates()

	fService, err := es.estateStorage.AddFacultyServiceSubscription(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) AddGuest(ctx context.Context, fService *model.Guest) (*model.Guest, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	fService, err := es.residentStorage.AddGuest(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) GetGuests(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Guest, error) {
	fsc, err := es.residentStorage.GetGuests(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) AddInvites(ctx context.Context, fService *model.Invite) (*model.Invite, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	// check resident status before allowing access
	if fService.ResidentID != "" {
		isBanned, _ := es.estateStorage.IsResidentBanned(ctx, fService.ResidentID)
		if isBanned {
			return nil, errors.ErrAuthError.New("resident is banned for a period of time")
		}
	} else if fService.HouseOwnerID != "" {
		isBanned, _ := es.estateStorage.IsOwnerBanned(ctx, fService.ResidentID)
		if isBanned {
			return nil, errors.ErrAuthError.New("resident is banned for a period of time")
		}
	}

	fService.InviteCode = fService.GenerateInviteCode()
	_, err := es.GetInviteByCode(ctx, fService.InviteCode)
	if err == nil {
		return nil, errors.ErrInternalServerError.New("unable to generate invite code, try again")
	}

	fService, err = es.residentStorage.AddInvites(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) GetInvites(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Invite, error) {
	fsc, err := es.residentStorage.GetInvites(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) GetInviteByCode(ctx context.Context, inviteCode string) (*model.Invite, error) {
	fsc, err := es.residentStorage.GetInviteByCode(ctx, inviteCode)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) VerifyCode(ctx context.Context, inv string) (*model.Invite, error) {
	invite, err := es.GetInviteByCode(ctx, inv)
	if err != nil {
		return nil, err
	}

	if invite.Status != constant.Active {
		return nil, errors.ErrInviteCodeInvalid.NewWithNoMessage()
	}

	now := time.Now()
	if !now.Before(invite.EndDate) {
		return nil, errors.ErrInviteCodeInvalid.New("invalid invite code")
	}

	return invite, nil
}

func (es *estate) UpdateInvite(ctx context.Context, invite *model.Invite) error {
	if err := invite.ValidateUpdate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return err
	}

	err := es.residentStorage.UpdateInvite(ctx, invite)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return err
	}
	return nil
}

func (es *estate) GetMonthlyReport(ctx context.Context, month time.Time, estateID string) (*model.MonthlyReport, error) {
	fsc, err := es.residentStorage.GetMonthlyReport(ctx, month, estateID)
	if err != nil {
		return nil, err
	}

	mrep, err := es.FetchHouseFeeSubscription(ctx, month, estateID)
	if err != nil {
		return nil, err
	}

	fsc.UnpaidAmount = mrep.Amount
	// mrep.Amount
	return fsc, nil
}

func (es *estate) FetchHouseFeeSubscription(ctx context.Context, month time.Time, estateId string) (*model.TenantReport, error) {
	fsc, err := es.estateStorage.FetchHouseFeeSubscription(ctx, month, estateId)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) FacultySubscriptionFeeReport(ctx context.Context, month time.Time, estateId string) ([]model.FacultyServiceMonthlyReport, error) {
	fsc, err := es.estateStorage.FacultySubscriptionFeeReport(ctx, month, estateId)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) InvitesReport(ctx context.Context, year int, estateId string) ([]model.InviteReport, error) {
	fsc, err := es.estateStorage.InvitesReport(ctx, year, estateId)
	if err != nil {
		return nil, err
	}
	return fsc, nil
}

func (es *estate) LogBooking(ctx context.Context, code string, fService *model.BookingLog) (*model.BookingLog, error) {
	if err := fService.Validate(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		es.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	inv, err := es.residentStorage.GetInviteByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	fService.InviteID = inv.InviteID

	now := time.Now()

	if !now.Before(inv.EndDate) && fService.Type == model.CheckOut {

		fService.IsDefaulted = true

		// check if ban is enabled
		config, err := es.estateStorage.GetEstateConfiguration(ctx, inv.House.EstateID)
		if err != nil {
			return nil, err
		}

		banStart := time.Now().Add(time.Hour * time.Duration(24))

		now := time.Now()
		totalGuests, err := es.residentStorage.GetGuestDefaultsPerWeek(ctx, inv.ResidentID, now)
		if err != nil {
			return nil, err
		}

		_, err = es.estateStorage.GetActiveDefault(ctx, inv.ResidentID)
		if err != nil {
			return nil, err
		}

		if err == nil {
			if config.IsInviteBanEnabled && totalGuests > config.MaxDefaultedInvites {
				_, _ = es.estateStorage.AddDefault(ctx, &model.Default{
					ResidentDefault: &model.ResidentDefault{
						ResidentID: inv.ResidentID,
						Reason:     "for payment",
					},
					IsPaymentOptionAvailable: true,
					PaymentAmount:            config.InviteDefaultPenaltyAmount,
					IsBanActive:              true,
					BanStartTime:             banStart,
					BanEndTime:               banStart.AddDate(0, int(config.BanMonths), int(config.BanDays)),
				})
			}
		}
	}

	fService, err = es.estateStorage.LogBooking(ctx, fService)
	if err != nil {
		es.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return fService, nil
}

func (es *estate) UpdateEstateConfiguration(ctx context.Context, fService *model.EstateConfiguration) (*model.EstateConfiguration, error) {
	if _, err := es.estateStorage.GetEstateConfiguration(ctx, fService.EstateID); err == nil {
		if err := fService.ValidateUpdate(); err != nil {
			err = errors.ErrInvalidInput.Wrap(err, "invalid input")
			es.logger.Info(ctx, "invalid input", zap.Error(err))
			return nil, err
		}

		fService, err := es.estateStorage.UpdateEstateConfiguration(ctx, fService)
		if err != nil {
			es.logger.Warn(ctx, err.Error())
			return nil, err
		}
		return fService, nil
	}
	return nil, errors.ErrNoRecordFound.New("estate configuration is not yet added")

}

func (es *estate) AddEstateConfiguration(ctx context.Context, fService *model.EstateConfiguration) (*model.EstateConfiguration, error) {
	if _, err := es.estateStorage.GetEstateConfiguration(ctx, fService.EstateID); err != nil {
		if err := fService.Validate(); err != nil {
			err = errors.ErrInvalidInput.Wrap(err, "invalid input")
			es.logger.Info(ctx, "invalid input", zap.Error(err))
			return nil, err
		}

		return es.estateStorage.AddEstateConfiguration(ctx, fService)
	}

	return nil, errors.ErrDataExists.New(errors.DataIsAlreadyRegistered)
}

func (es *estate) GetEstateConfiguration(ctx context.Context, estateID string) (*model.EstateConfiguration, error) {
	return es.estateStorage.GetEstateConfiguration(ctx, estateID)
}

func (es *estate) BanHouseOwnersForUnpaidServices(ctx context.Context) error {
	return es.estateStorage.BanHouseOwnersForUnpaidServices(ctx)
}
