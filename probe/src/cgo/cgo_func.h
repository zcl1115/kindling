//
// Created by jundi zhou on 2022/6/1.
//

#ifndef SYSDIG_CGO_FUNC_H
#define SYSDIG_CGO_FUNC_H

#ifdef __cplusplus
extern "C" {
#endif
int runForGo();
int getKindlingEvent(void** kindlingEvent);
void subEventForGo(char* eventName, char* category, void* params);
int startProfile();
int stopProfile();
void startProfileDebug(int pid, int tid);
void stopProfileDebug();
#ifdef __cplusplus
}
#endif

#endif  // SYSDIG_CGO_FUNC_H
