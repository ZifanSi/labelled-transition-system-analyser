@echo off
echo Launching LTSA with Sample folder...

:: Set LTSA path
set LTSA_JAR=ltsa.jar

:: Set working directory (adjust if needed)
cd /d "%~dp0Sample"

:: Launch LTSA (just opens the app in Sample dir)
java -jar "..\%LTSA_JAR%"

pause
