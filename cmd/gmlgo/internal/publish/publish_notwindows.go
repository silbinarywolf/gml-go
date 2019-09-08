// +build !windows

package publish

import "golang.org/x/xerrors"

func compile(dir string, distFolder string, args []string) error {
	if err := compileWeb(dir, distFolder, args); err != nil {
		return xerrors.Errorf("error compiling web: %w", err)
	}
	if err := compileWindows(dir, distFolder, args); err != nil {
		return err
	}
	if err := compileLinux(dir, distFolder, args); err != nil {
		return err
	}
	if err := compileMac(dir, distFolder, args); err != nil {
		return err
	}
	return nil
}
