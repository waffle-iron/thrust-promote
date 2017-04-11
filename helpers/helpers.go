package helpers

import (
    "os"
    "log"
    "os/exec"
    "strings"
    "path/filepath"
)

func RemoveFileExt(filename string) string {
    return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func GetAudioLength(filename string) int {
    var stdErr bytes.Buffer
    cmdString := []string{"-i", filename, "2>&1", "|", "grep Duration", "|",
       "sed", `'s/Duration: \(.*\), start/\1/g'`}
    cmd := exec.Command("ffmpeg", cmdString...)
    cmd.Stderr = &stdErr
    if err := cmd.Start(); err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to start: %v", err)
        return false, err
    }

    err = cmd.Wait() 
    if err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to finish: %v", err)
        return false, err
    }
}

func ConvertAudioCommand(filename string, targetFilename string) (string, error) {
    var stdErr bytes.Buffer
    cmd := exec.Command("ffmpeg", "-i", filename, targetFilename)
    cmd.Stderr = &stdErr
    if err := cmd.Start(); err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to start: %v", err)
        return false, err
    }

    err = cmd.Wait() 
    if err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to finish: %v", err)
        return false, err
    }

    return targetFilename, nil
}

func ConvertVideoCommand(filename string, imageFilename string, videoTargetFilename string) string {
    audioLength := GetAudioLength(filename)
    // cmd := `ffmpeg -y -loop 1 -f image2 -i #{@image_file} \
    //      -i "#{@audio_file}" -c:v libx264 -c:a aac -strict experimental \
    //      -b:a 192k -t #{audio_length} #{@video_file}`
    cmdString := []string{"-y", "-loop", 1, "-f", "image2", "-i",
        imageFilename, "-i", filename, "-c:v", "libx264", "-c:a", "aac",
        "-strict", "experimental", "-b:a", "193k", "-t", audioLength,
        videoTargetFilename}

    var stdErr bytes.Buffer
    cmd := exec.Command("ffmpeg", cmdString...)
    cmd.Stderr = &stdErr
    if err := cmd.Start(); err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to start: %v", err)
        return false, err
    }

    err = cmd.Wait() 
    if err != nil {
        log.Println(stdErr.String())
        log.Fatalf("Command failed to finish: %v", err)
        return false, err
    }

    return videoTargetFilename, nil
}

func RemoveFiles(filenames []string) {
    for _, filename := range filenames {
        os.Remove(filename)
    }
}