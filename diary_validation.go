package thenovadiary

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

type validationCheck func(cfg *DiaryConfig) string

var (
	// --------------------------------------
	// Define the hard coded checks for the
	// diary configuration here
	//
	// TODO Use reflection to pre-populate all
	// functions that start with "v_"
	validationChecks = []validationCheck{
		v_emptyTwitterPass,
		v_emptyTwitterUser,
		//v_actionStringToAction,
		v_emptyPhotoprismUser,
		v_emptyPhotoprismPass,
		v_emptyPhotoprismConn,
	}
)

func ValidateConfig(cfg *DiaryConfig) error {
	var strerr string
	for _, v := range validationChecks {
		e := v(cfg)
		if e != "" {
			strerr = fmt.Sprintf("%s%s", strerr, e)
		}
	}
	if strerr != "" {
		return fmt.Errorf("Errors during validation: %s", strerr)
	}
	// We know we have a valid config
	cfg.validated = true
	return nil
}

//func v_actionStringToAction(cfg *DiaryConfig) string {
//	spl := strings.Split(cfg.ActionString, ",")
//	if len(spl) <= 0 {
//		return ferr(fmt.Sprintf("Invalid ActionString %s", cfg.ActionString))
//	}
//	for _, a := range spl {
//		pKey := strings.Replace(a, " ", "", -1)
//		if f, ok := ActionMap[pKey]; ok {
//			logger.Always("Mapping action [%s]", pKey)
//			cfg.Actions = append(cfg.Actions, f)
//		} else {
//			return ferr(fmt.Sprintf("Invalid action: %s", pKey))
//		}
//	}
//	return ""
//}

func v_emptyTwitterPass(cfg *DiaryConfig) string {
	if cfg.TwitterPass == "" {
		cfg.TwitterPass = os.Getenv("DIARY_TWITTERPASS")
	}
	if cfg.TwitterPass == "" {
		return ferr("Empty TwitterPass")
	}
	if len(cfg.TwitterPass) < 3 {
		return ferr("TwitterPass < 3 chars")
	}
	return ""
}

func v_emptyTwitterUser(cfg *DiaryConfig) string {
	if cfg.TwitterUser == "" {
		cfg.TwitterUser = os.Getenv("DIARY_TWITTERUSER")
	}
	if cfg.TwitterUser == "" {
		return ferr("Empty TwitterUser")
	}
	if len(cfg.TwitterPass) < 3 {
		return ferr("TwitterUser < 3 chars")
	}
	return ""
}

func v_emptyPhotoprismUser(cfg *DiaryConfig) string {
	if cfg.PhotoprismUser == "" {
		cfg.PhotoprismUser = os.Getenv("DIARY_PHOTOPRISMUSER")
	}
	if cfg.PhotoprismUser == "" {
		return ferr("Empty PhotoprismUser")
	}
	if len(cfg.PhotoprismUser) < 3 {
		return ferr("PhotoprismUser < 3 chars")
	}
	return ""
}

func v_emptyPhotoprismPass(cfg *DiaryConfig) string {
	if cfg.PhotoprismPass == "" {
		cfg.PhotoprismPass = os.Getenv("DIARY_PHOTOPRISMPASS")
	}
	if cfg.PhotoprismPass == "" {
		return ferr("Empty PhotoprismPass")
	}
	if len(cfg.PhotoprismPass) < 3 {
		return ferr("PhotoprismPass < 3 chars")
	}
	return ""
}

func v_emptyPhotoprismConn(cfg *DiaryConfig) string {
	if cfg.PhotoprismConn == "" {
		cfg.PhotoprismConn = os.Getenv("DIARY_PHOTOPRISMCONN")
	}
	if cfg.PhotoprismConn == "" {
		return ferr("Empty PhotoprismConn")
	}
	if len(cfg.PhotoprismConn) < 3 {
		return ferr("PhotoprismConn < 3 chars")
	}
	return ""
}

func ferr(str string) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next() // Ignoring next bool (This is error safe)
	fqn := frame.Func.Name()
	spl := strings.Split(fqn, ".")
	if len(spl) < 3 {
		return fmt.Sprintf("[%s] %s ", frame.Func.Name(), str)
	}
	return fmt.Sprintf("[%s] %s ", spl[2], str)
}
