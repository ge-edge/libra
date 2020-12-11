# Project Libra
# author: apolo.yasuda@ge.com 
#
cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-wzlib)
find_library(golang REQUIRED)

set(TARGET lib)
set(GOPATH ${CMAKE_CURRENT_BINARY_DIR})

set(SRC crypto.go)
set(LIB crypto.so)
set(INCLUDE crypto.h)

add_custom_command(OUTPUT ./../lib/${LIB}
  DEPENDS ${SRCS}
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
  COMMAND env GOPATH=${GOPATH} go build -buildmode=c-shared
  -o "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
  ${CMAKE_GO_FLAGS} ./...
  COMMENT "Building Go library")

add_custom_target(${TARGET} DEPENDS ${LIB} ${HEADER})

add_library(wzlib STATIC IMPORTED GLOBAL)
add_dependencies(wzlib ${TARGET})
set_target_properties(wzlib
  PROPERTIES
  IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
  INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR})
