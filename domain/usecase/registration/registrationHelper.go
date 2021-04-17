package registration

func modifyUser(udi dataservice.UserDataInterface, user *model.User) error {
	err := user.ValidatePersisted()
	if err != nil {
		return errors.Wrap(err, "user validation failed")
	}
	rowsAffected, err := udi.Update(user)
	if err != nil {
		return errors.Wrap(err, "")
	}

	if rowsAffected != 1 {
		return errors.Wrap("Modify user failed. rows affected is " + strconv.Itoa(int(rowsAffected)))
	}

	return nil
}

func unregisterUser(udi dataservice.UserDataInterface, username string) error {
	affected, err := udi.Remove(username)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if affected == 0 {
		errStr := "UnregisterUser failed. No such user " + username
		return errors.New(errStr)
	}

	if affected != 1 {
		errStr := "UnregisterUser failed. Number of users unregistered are " + strconv.Itoa(int(affected))
		return errors.New(errStr)
	}
	return nil
}

// ModifyAndUnregister the business function will be wrapped inside a transaction or a non-transaction function
// It needs to be writtedn in a way that every error will be returne so it can be caught by TxEnd() function,
// which will handle commit and rollback
func ModifyAndUnregister(udi dataservice.UserDataInterface, user *model.User) error {
	err := modifyUser(udi, user)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = unregisterUser(udi, user.Name)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
