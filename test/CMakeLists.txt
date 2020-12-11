cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

add_library(common-lib SHARED IMPORTED GLOBAL)
set_target_properties(
  common-lib
  PROPERTIES
  IMPORTED_LOCATION /home/pi/ged/libra/lib/crypto.so
)

add_executable(libra test_libra.cpp)

target_link_libraries(libra PRIVATE common Threads::Threads common-lib)

install(TARGETS libra DESTINATION bin)
