# Project Libra
# author: apolo.yasuda@ge.com 
#
cmake_minimum_required(VERSION 3.8)
project(libra)

set(CMAKE_CXX_STANDARD 20)

set(CMAKE_INSTALL_PREFIX install)

set(PROJECT_INCLUDE_DIR include)
set(PROJECT_SOURCE_DIR src)
set(COMMON_INCLUDE_DIR common/include)
set(WZLIB_GO_DIR wzlib)

include_directories(${PROJECT_INCLUDE_DIR})
include_directories(${PROJECT_SOURCE_DIR})
include_directories(${COMMON_INCLUDE_DIR})
include_directories(${WZLIB_GO_DIR})
 
# cgo library folder code
add_subdirectory(wzlib)

# Common code, packaged as a library
add_subdirectory(common)

# Sub-projects, packaged as libraries
add_subdirectory(src)

# Tests for Sub-projects, as execiutables
add_subdirectory(test)
