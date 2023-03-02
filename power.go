package main

func PowerUp(lastProfile Profile, upTimes int) Profile {
	newProfile := Profile{
		Name:   lastProfile.Name,
		Health: lastProfile.Health + (lastProfile.Health * upTimes),
		Power:  lastProfile.Power + (lastProfile.Power * upTimes),
		Exp:    lastProfile.Exp + (lastProfile.Exp * upTimes),
	}
	return newProfile
}
