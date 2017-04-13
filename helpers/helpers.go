package helpers

import (
    "os"
    "log"
    "os/exec"
    "bytes"
    "strings"
    "strconv"
    "path/filepath"
)

func RemoveFileExt(filename string) string {
    return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func RemoveFiles(filenames []string) {
    for _, filename := range filenames {
        os.Remove(filename)
    }
}

func GetAudioLength(filename string) (float64, error) {
    var stdErr bytes.Buffer
    var stdOut bytes.Buffer
    cmd := exec.Command("soxi", "-D", filename)
    cmd.Stderr = &stdErr
    cmd.Stdout = &stdOut
    if err := cmd.Start(); err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to start: %v", err)
        return 0, err
    }

    err := cmd.Wait() 
    if err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to finish: %v", err)
        return 0, err
    }

    resultStr := strings.Trim(stdOut.String(), "\r\n")
    duration, err := strconv.ParseFloat(resultStr, 64)

    if err != nil {
        return 0, err
    }

    return duration, nil
}

func ConvertAudioCommand(filename string, targetFilename string) error {
    var stdErr bytes.Buffer
    cmd := exec.Command("ffmpeg", "-i", filename, targetFilename)
    cmd.Stderr = &stdErr
    if err := cmd.Start(); err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to start: %v", err)
        return err
    }

    err := cmd.Wait() 
    if err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to finish: %v", err)
        return err
    }

    return nil
}

func ConvertVideoCommand(filename string, imageFilename string, videoTargetFilename string) error {
    audioLength, err := GetAudioLength(filename)
    if err != nil {
        return err
    }

    audioLengthStr := strconv.FormatFloat(audioLength, 'f', -1, 64)
    cmdString := []string{"-y", "-loop", "1", "-f", "image2", "-i",
        imageFilename, "-i", filename, "-c:v", "libx264", "-c:a", "aac",
        "-strict", "experimental", "-b:a", "193k", "-t", audioLengthStr,
        videoTargetFilename}

    var stdErr bytes.Buffer
    cmd := exec.Command("ffmpeg", cmdString...)
    cmd.Stderr = &stdErr
    if err := cmd.Start(); err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to start: %v", err)
        return err
    }

    err = cmd.Wait() 
    if err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to finish: %v", err)
        return err
    }

    return nil
}