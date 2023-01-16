package app

import "fmt"

func devFilesConcats(f string) (fc []string) {

	// File path concat
	fcc := fmt.Sprintf("./backups/dev-env/%s/backup-dbchat-%s", f, f)
	fce := fmt.Sprintf("./backups/dev-env/%s/backup-dbeco-%s", f, f)
	fcp := fmt.Sprintf("./backups/dev-env/%s/backup-dbpay-%s", f, f)

	fc = append(fc, fcc, fce, fcp)

	return
}
