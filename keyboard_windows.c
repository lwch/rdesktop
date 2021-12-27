#include "keyboard_windows.h"

void keyboard_toggle(uint32_t code, bool down) {
    int scan = MapVirtualKey(code & 0xff, MAPVK_VK_TO_VSC);
    INPUT keyInput;
    keyInput.type = INPUT_KEYBOARD;
    keyInput.ki.wVk = code;
    keyInput.ki.wScan = scan;
    keyInput.ki.dwFlags = down ? 0 : KEYEVENTF_KEYUP;
    keyInput.ki.time = 0;
    keyInput.ki.dwExtraInfo = 0;
    SendInput(1, &keyInput, sizeof(keyInput));
}