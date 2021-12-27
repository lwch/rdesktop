#include "mouse_windows.h"

void mouse_move(uint32_t x, uint32_t y) {
    #define MOUSE_COORD_TO_ABS(coord, width_or_height) ( \
        ((65536 * coord) / width_or_height) + (coord < 0 ? -1 : 1))

    INPUT mouseInput;
    mouseInput.type = INPUT_MOUSE;
    mouseInput.mi.dx = MOUSE_COORD_TO_ABS(x, GetSystemMetrics(SM_CXSCREEN));
    mouseInput.mi.dy = MOUSE_COORD_TO_ABS(y, GetSystemMetrics(SM_CYSCREEN));
    mouseInput.mi.dwFlags = MOUSEEVENTF_ABSOLUTE | MOUSEEVENTF_MOVE;
    mouseInput.mi.time = 0;
    mouseInput.mi.dwExtraInfo = 0;
    mouseInput.mi.mouseData = 0;
    SendInput(1, &mouseInput, sizeof(mouseInput));
}

void mouse_toggle(uint32_t button, bool down) {
    INPUT mouseInput;
    mouseInput.type = INPUT_MOUSE;
    mouseInput.mi.dx = 0;
    mouseInput.mi.dy = 0;
    mouseInput.mi.time = 0;
    mouseInput.mi.dwExtraInfo = 0;
    mouseInput.mi.mouseData = 0;
    switch (button) {
    case 0:  // left
        mouseInput.mi.dwFlags = down ? MOUSEEVENTF_LEFTDOWN : MOUSEEVENTF_LEFTUP;
        break;
    case 1:  // right
        mouseInput.mi.dwFlags = down ? MOUSEEVENTF_RIGHTDOWN : MOUSEEVENTF_RIGHTUP;
        break;
    default: // middle
        mouseInput.mi.dwFlags = down ? MOUSEEVENTF_MIDDLEDOWN: MOUSEEVENTF_MIDDLEUP;
        break;
    }
    SendInput(1, &mouseInput, sizeof(mouseInput));
}

void scroll(uint32_t x, uint32_t y) {
    INPUT mouseScrollInputH;
    INPUT mouseScrollInputV;

    mouseScrollInputH.type = INPUT_MOUSE;
    mouseScrollInputH.mi.dx = 0;
    mouseScrollInputH.mi.dy = 0;
    mouseScrollInputH.mi.dwFlags = MOUSEEVENTF_WHEEL;
    mouseScrollInputH.mi.time = 0;
    mouseScrollInputH.mi.dwExtraInfo = 0;
    mouseScrollInputH.mi.mouseData = WHEEL_DELTA * x;

    mouseScrollInputV.type = INPUT_MOUSE;
    mouseScrollInputV.mi.dx = 0;
    mouseScrollInputV.mi.dy = 0;
    mouseScrollInputV.mi.dwFlags = MOUSEEVENTF_WHEEL;
    mouseScrollInputV.mi.time = 0;
    mouseScrollInputV.mi.dwExtraInfo = 0;
    mouseScrollInputV.mi.mouseData = WHEEL_DELTA * y;

    SendInput(1, &mouseScrollInputH, sizeof(mouseScrollInputH));
    SendInput(1, &mouseScrollInputV, sizeof(mouseScrollInputV));
}