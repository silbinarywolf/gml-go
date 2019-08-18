// +build windows

package publish

func compile(dir string, distFolder string, args []string) error {
	if err := compileWeb(dir, distFolder, args); err != nil {
		return err
	}
	if err := compileWindows(dir, distFolder, args); err != nil {
		return err
	}
	// NOTE(Jake): 2019-08-16
	// Skip these as I get the following on Windows 10:
	// - cannot load github.com/go-gl/glfw/v3.2/glfw: no Go source files
	//
	// This is most likely due to a lack of the package but it
	// could also fail if a publisher doesn't have CGo.
	//if err := compileLinux(dir, distFolder, args); err != nil {
	//	return err
	//}
	//if err := compileMac(dir, distFolder, args); err != nil {
	//	return err
	//}
	return nil
}
