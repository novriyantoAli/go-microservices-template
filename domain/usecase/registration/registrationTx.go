package registration

// RegistrationTxUseCase implements RegistrationTxUseCaseInterface.
// It has UserDataInterface, which can be used to access persistence layer
type RegistrationTxUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

// ModifyAndUnregisterWithTx is a use case with transaction
func (rtuc *RegistrationTxUseCase) ModifyAndUnregisterWithTx(user *model.User) error {

	udi := rtuc.UserDataInterface
	return udi.EnableTx(func() error {
		// wrap the business function inside the TxEnd function
		return ModifyAndUnregister(udi, user)
	})
}
