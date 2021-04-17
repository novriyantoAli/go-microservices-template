// Package registrationTx represents the concrete implementation of RegistrationUseCaseInterface and
// RegistrationTxUseCaseInterface interface.
// Because the same business function can be created to support both transaction and non-transaction,
// a shared business function is created in a helper file, then we can wrap that function with transaction
// or non-transaction.

package registration

// RegistrationUseCase implements RegistrationUseCaseInterface.
// It has UserDataInterface, which can be used to access persistence layer
// TxDataInterface is needed to support transaction
type RegistrationUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

// RegisterUser ...
func (ruc *RegistrationUseCase) RegisterUser(user *model.User) (*model.User, error) {
	err := user.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "user validation failed")
	}
	isDup, err := ruc.isDuplicate(user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if isDup {
		return nil, errors.New("duplicate user for " + user.Name)
	}
	resultUser, err := ruc.UserDataInterface.Insert(user)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return resultUser, nil
}

// ModifyUser ...
func (ruc *RegistrationUseCase) ModifyUser(user *model.User) error {
	return modifyUser(ruc.UserDataInterface, user)
}

func (ruc *RegistrationUseCase) isDuplicate(name string) (bool, error) {
	user, err := ruc.UserDataInterface.FindByName(name)
	//logger.Log.Debug("isDuplicate() user:", user)
	if err != nil {
		return false, errors.Wrap(err, "")
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}

// UnregisterUser ...
func (ruc *RegistrationUseCase) UnregisterUser(username string) error {
	return unregisterUser(ruc.UserDataInterface, username)
}

// ModifyAndUnregister is a use case without transaction
func (ruc *RegistrationUseCase) ModifyAndUnregister(user *model.User) error {
	return ModifyAndUnregister(ruc.UserDataInterface, user)
}
