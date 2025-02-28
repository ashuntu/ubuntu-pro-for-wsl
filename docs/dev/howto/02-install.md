# How to install Ubuntu Pro for WSL

This guide will show you how to install Ubuntu Pro for WSL for local development and testing.

## Install Ubuntu Pro for WSL from scratch

### Requirements

- A Windows machine with access to the internet
- Appx from the Microsoft Store:
  - Windows Subsystem For Linux
  - Either Ubuntu, Ubuntu 22.04, or Ubuntu (Preview)
- The Windows Subsystem for Windows optional feature enabled

### Download
<!-- TODO: Update when we change were artifacts are hosted -->
1. Go to the [repository actions page](https://github.com/canonical/ubuntu-pro-for-wsl/actions/workflows/qa-azure.yaml?query=branch%3Amain+).
2. Click the latest successful workflow run.
3. Scroll down past any warnings or errors, until you reach the Artifacts section.
4. Download:
    - Windows agent:    UbuntuProForWSL+...-production
    - wsl-pro-service:  Wsl-pro-service_...

Notice that, for the step above, there is also an alternative version of the MSIX bundle enabled for end-to-end testing. Most likely, that's not what you want to download.

### Install the Windows agent

This is the Windows-side agent that manages the distros.

1. Uninstall Ubuntu Pro for WSL if you had installed previously:

    ```powershell
    Get-AppxPackage -Name ubuntu-pro-for-wsl | Remove-AppxPackage
    ```

2. Follow the download steps to download UbuntuProForWSL
3. Unzip the artefact
4. Find the certificate inside. Install it into `Local Machine/Trusted people`.
5. Double click on the MSIX bundle and complete the installation.
6. The Firewall may ask for an exception. Allow it.
7. The GUI should show up. You’re done.

### Install the WSL Pro Service

This is the Linux-side component that talks to the agent. Choose one or more distros Jammy or greater, and follow the instructions.

1. Uninstall the WSL-Pro-Service from your distro if you had it installed previously:

    ```bash
    sudo apt remove wsl-pro-service
    ```

2. Follow the download steps to download the WSL-Pro-Service.
3. Unzip the artifact.
4. Navigate to the unzipped directory containing the .deb file. Here is a possible path:

    ```bash
    cd /mnt/c/Users/WINDOWS-USER/Downloads/wsl-pro-service_*
    ```

5. Install the deb package.

    ```bash
    sudo apt install ./wsl-pro-service_*.deb
    ```

6. Ensure it works via systemd:

    ```bash
    systemctl status wsl-pro.service
    ```

## Reset Ubuntu Pro for WSL back to factory settings

You can reset Ubuntu Pro for WSL to factory settings following these steps:

1. Uninstall the package and shut down WSL:

    ```powershell
    Get-AppxPackage -Name "CanonicalGroupLimited.UbuntuProForWSL" | Remove-AppxPackage`
    wsl --shutdown
    ```

2. Remove registry key `HKEY_CURRENT_USER\Software\Canonical\UbuntuPro`.
3. Install the package again (see the section on [how to install](./02-install.md)).
4. You're done. Next time you start the GUI it'll be like a fresh install.
