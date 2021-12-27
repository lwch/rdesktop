#ifndef _MOUSE_WINDOWS_H_
#define _MOUSE_WINDOWS_H_

#include <windows.h>
#include <stdint.h>
#include <stdbool.h>

void mouse_move(uint32_t x, uint32_t y);
void mouse_toggle(uint32_t button, bool down);
void scroll(uint32_t x, uint32_t y);

#endif