# Beacon
The code for the beacon is basically [the ble_app_beacon](https://infocenter.nordicsemi.com/index.jsp?topic=%2Fcom.nordic.infocenter.sdk5.v15.0.0%2Fble_sdk_app_beacon.html) example from Nordic. The only change made in this example is that the `APP_MINOR_VALUE` and `APP_MINOR_VALUE` is changed to **0xdeadc0de** in order for the gateway to filter out only the BLE data meant for the hackathon. The idea is that each of the groups can use a different `APP_BEACON_UUID` in order to identify their device.


## nRF52 firmware development setup
Developing code for Nordic's SoCs can be done with many different IDEs with various levels of complexity. My definitive favorite and recommendation is to simply use a text editor of your choice, and GNU Make together with GCC to build and flash the code. This guide will show how to setup a development environment using the text editor VS code and a makefile from Nordic that we will modify slightly. 

### Installation
* Install VS code, and its C/C++ extension
* Download the [nR52 SDK](https://developer.nordicsemi.com/nRF5_SDK/nRF5_SDK_v15.x.x/) (This example uses version 15.0.0). Extract the SDK to some appropriate location
* Install [GNU Arm Embedded Toolchain](https://developer.arm.com/tools-and-software/open-source-software/developer-tools/gnu-toolchain/gnu-rm/downloads)
* Install [nRF Command Line Tools](https://www.nordicsemi.com/?sc_itemid=%7B56868165-9553-444D-AA57-15BDE1BF6B49%7D)
* In your SDK folder, find the file *components/toolchain/gcc/Makefile.posix* (or Makefile.windows if you are on windows) and update it so that `GNU_INSTALL_ROOT` points to the bin folder in the location where you installed GNU Arm Embedded Toolchain.
* You should now be able to run `make` in the subfolder *pca10040/s132/armgcc* (pca number depending on your board) for the examples in the SDK folder in order to compile the project. Running `make flash` will flash the connected board.

### VS code setup
VS code is a neat editor, that provides a pleasant experience when coding. However, if you at this point open the ble_app_beacon example in VS code, you'll get a bunch of red squiggly lines and warnings. In order to get proper [intelliSense](https://code.visualstudio.com/docs/editor/intellisense), we need to help the editor understand where to find the source files of the SDK. Specifically, if you press F1 in VS code and search for *C/C++: Edit Configurations (JSON)* you'll get a file that has the settings we are looking for. We need to edit `includePath`, `defines` and `compilerPath`. The `compilerPath` should be set to the full path of arm-none-eabi-gcc, which is one of the things you installed in the last step. The `IncludePath` and the `defines` needs to be updated with the equivalent values from the Makefile, unfortunately the corresponding values in the Makefile are on a entirely different format from what we need to have in the JSON file. In order to help us, we can write some phony makefile targets that will print the values in the format we want:

```
.PHONY: print_includepath print_defines
print_includepath: 
	@$(foreach _FOLDER, $(realpath $(INC_FOLDERS)), echo \"$(_FOLDER)\",;)

print_defines:
	@$(foreach _FLAG, $(subst -D,,$(filter -D%,$(CFLAGS))), echo \"$(_FLAG)\",;)
```
Add the above code to the bottom of your Makefile, and run `make print_includepath` and `make print_defines`. The output can then be pasted directly into the JSON file!


### Workflow
When developing with this setup, you are simply using a text editor instead of an IDE. This means that you are left to control things such as which files to include and build, as well as the configuration of the SDK. Let's say that you need to use the nrfx_spi driver in your project. To do this, you'll have to add the file *$(SDK_ROOT)/modules/nrfx/drivers/src/nrfx_spi.c* to the `SRC_FILES` variable in the makefile. Also, if not already added you'll need to add *$(SDK_ROOT)/modules/nrfx/drivers/include* to the `INC_FOLDERS` variable. If you added a new folder to `INC_FOLDERS`, you'll also need to update the `includePath` in *c_cpp_properties.json* in order to get proper intelliSense. The configuratuin file *pca1400/s132/config/sdk_config.h* is also left for you to update when enabling new modules.