Add-Type -AssemblyName system.windows.forms

function Create-StorageDirectory($Directory_Name){
try {
if ((Test-Path -Path "C:\$Directory_Name") -eq $false){
New-Item -Path c:\ -ItemType directory -Name $Directory_Name -Confirm:$false -Force -ErrorAction Stop
}
return "Created directory $Directory_Name"
} catch {
return "Can't create directory $Directory_Name"
}
}
function Download-FromGit($Destination, $Git_Location){
try{
Invoke-WebRequest -Uri $Git_Location -OutFile $Destination -ErrorAction Stop
return "Succes"
} catch {
return "Failed"
}
}
try{
$Git_Location = "https://github.com/M8tan/Evringlish/releases/latest/download"
Create-StorageDirectory -Directory_Name "Evringlish"
Download-FromGit -Destination "C:\Evringlish\EvringlishApp.exe" -Git_Location "$Git_Location/EvringlishApp.exe"
Download-FromGit -Destination "C:\Evringlish\EvringlishFirst.exe" -Git_Location "$Git_Location/EvringlishApp.exe"
Download-FromGit -Destination "C:\Evringlish\readme.txt" -Git_Location "$Git_Location/readme.txt"
Download-FromGit -Destination "C:\Evringlish\Updates.txt" -Git_Location "$Git_Location/Updates.txt"
Download-FromGit -Destination "C:\Evringlish\Updater.bat" -Git_Location "$Git_Location/Updater.bat"
New-Item -Path "C:\Evringlish" -ItemType file -Name EvringlishUpdate.bat -Value '
@echo off
title Get newest Evringlish version
powershell.exe -NoProfile -ExecutionPolicy Bypass -Command "& {Invoke-WebRequest -Uri "https://github.com/M8tan/Evringlish/releases/latest/download/EvringlishApp.exe" -OutFile "C:\Evringlish\EvringlishApp.exe"}"
powershell.exe -NoProfile -ExecutionPolicy Bypass -Command "& {Invoke-WebRequest -Uri "https://github.com/M8tan/Evringlish/releases/latest/download/Updates.txt" -OutFile "C:\Evringlish\Updates.txt"}"
powershell.exe -NoProfile -ExecutionPolicy Bypass -Command "& {Invoke-WebRequest -Uri "https://github.com/M8tan/Evringlish/releases/latest/download/Updater.bat" -OutFile "C:\Evringlish\Updater.bat"}"
powershell.exe -NoProfile -ExecutionPolicy Bypass -Command "& {Invoke-WebRequest -Uri "https://github.com/M8tan/Evringlish/releases/latest/download/readme.txt" -OutFile "C:\Evringlish\readme.txt"}"
start c:\Evringlish\Updater.bat
echo Done updating!
timeout /t 3 /nobreak>nul
exit
' -Confirm:$false
New-Item -Path "C:\Evringlish" -ItemType file -Name Evringlish.bat -Value '
@echo off
title Run current Evringlish version
start c:\Evringlish\EvringlishApp.exe
' -Confirm:$false
New-Item -Path "C:\Evringlish" -ItemType file -Name EvringlishFull.bat -Value '
@echo off
title Run newest Evringlish version
powershell.exe -NoProfile -ExecutionPolicy Bypass -Command "& {Invoke-WebRequest -Uri "https://github.com/M8tan/Evringlish/releases/latest/download/EvringlishApp.exe" -OutFile "C:\Evringlish\EvringlishApp.exe"}"
start c:\Evringlish\EvringlishApp.exe
' -Confirm:$false
$envPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($envPath -notlike "*C:\Evringlish*") {
    [Environment]::SetEnvironmentVariable("Path", "$envPath;C:\Evringlish", "User")
}
$Trigger = New-ScheduledTaskTrigger -Daily -At 10:10am
$Action = New-ScheduledTaskAction -Execute C:\Evringlish\EvringlishUpdate.bat
$Task = New-ScheduledTask -Action $Action -Trigger $Trigger -Description "Update Evringlishs readme, application, updater and updates files"
Register-ScheduledTask -InputObject $Task -TaskName "Evringlish" -Force
([System.Windows.Forms.MessageBox]::Show("Evringlish installed succesfully! " + [environment]::NewLine + "Details can be found @ C:\Evringlish", "Done", [System.Windows.Forms.MessageBoxButtons]::OK, [System.Windows.Forms.MessageBoxIcon]::Information))

} catch {
([System.Windows.Forms.MessageBox]::Show("Unknown error occurred", "Error", [System.Windows.Forms.MessageBoxButtons]::OK, [System.Windows.Forms.MessageBoxIcon]::Error))
}
