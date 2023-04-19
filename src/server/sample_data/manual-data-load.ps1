
<#
LOAD SAMPLE DATA INTO APP DB
#>

#  SET COMMAND LINE ARGUMENTS
param (
    [Parameter(Mandatory = $false)]
    [string[]] $objects
)

#  Whitespace padding for console log readability
Write-Output "`n"

#  FLAG DEFINITION (DETERMINE IF ONLY SUBSET OF OBJECTS IS BEING LOADED)
if ($objects.Length -eq 0) {
    $all_objects = $True
} else {
    $all_objects = $False
}

#  UPDATE VARIABLES WITH APPROPRIATE VALUES
$DB_HOST="localhost"
$DB_PORT=8420

$OBJ_KEYS = "User" `
            , "Business" `
            , "Service" `
            , "Appointment" `
            , "Invoice" `
            

$API_ENDPOINTS = @{
                    User = "register" `
                    ; Business = "business" `
                    ; Service = "service"
                    ; Appointment = "appointment"
                    ; Invoice = "invoice" `
                  }

$FILE_NAMES = @{ `
                    User = "sample-users.json" `
                    ; Business = "sample-businesses.json" `
                    ; Service = "sample-services.json" `
                    ; Appointment = "sample-appointments.json" `
                    ; Invoice = "sample-invoices.json" `
                }

$PS_SCRIPT_FULL_PATH = $MyInvocation.MyCommand.Path
$PS_SCRIPT_BASE_PATH = Split-Path -Path "$PS_SCRIPT_FULL_PATH"

$loaded_objects = [System.Collections.ArrayList]::new()
# Iterate over keys and execute curl commands for each endpoint/sample data file
foreach($obj_key in $OBJ_KEYS) {

    if ($all_objects -or ${objects}.Contains($obj_key)) {

        Write-Output "STATUS:`t`tLoading sample $obj_key data."

        $api_endpoint = $API_ENDPOINTS[$obj_key]
        $sample_data_file_name = $FILE_NAMES[$obj_key]

        foreach($line in Get-Content $PS_SCRIPT_BASE_PATH\$sample_data_file_name) {
            Invoke-RestMethod -Uri http://${DB_HOST}:${DB_PORT}/$api_endpoint -Method POST -Body $line -ContentType "application/json" 
        }

        [void]$loaded_objects.Add($obj_key)
    }

}

#  Alert user if objects specified in CLI argument were not found in list of defined/supported objects
if (!$all_objects) {
    foreach($obj in $objects) {
        if (!${loaded_objects}.Contains($obj)) {
            Write-Output "ERROR:`t`tObject specified in -objects argument was not found in the list of supported objects for sample loads  --  $obj"
        }
    }
}

#  Whitespace padding for console log readability
Write-Output "`n"