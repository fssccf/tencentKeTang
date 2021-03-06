package ffmpeg

import (
	"bytes"
	"github.com/pkg/errors"
	"os/exec"
)

func (f *Ffmpeg) probe(vodUrl string) (ret string, err error) {
	probeArgs := []string{"-show_format", "-show_streams",
		"-of", "json"}
	args := append(probeArgs, vodUrl)
	cmd := exec.Command(f.ffprobeExec, args...)
	buf := bytes.NewBuffer(nil)
	errBuf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = errBuf
	err = cmd.Run()
	if err != nil {
		return "", errors.Wrapf(err, "exec.Run: %s", errBuf.String())
	}
	return buf.String(), nil
}

func (f *Ffmpeg) checkProbe() (err error) {
	cmd := exec.Command(f.ffprobeExec, "-version")
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, "exec.Run")
	}
	return nil
}
