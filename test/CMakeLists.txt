cmake_minimum_required(VERSION 3.13)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)
        
#set_source_files_properties(test_libra.cpp PROPERTIES LANGUAGE CXX)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)
#link_directories(${CMAKE_CURRENT_LIST_DIR})

message( "***** CMAKE_CURRENT_LIST_DIR: ${CMAKE_CURRENT_LIST_DIR}")

#find_library(CRYPTO_LIB
#         NAMES 
#            crypto.so
#         PATHS 
#            /home/runner/work/libra/libra/lib/linux
#)

#set(legacy_link ${CMAKE_CURRENT_LIST_DIR}/libcrypto.so)
#set(legacy_target /home/runner/work/libra/libra/lib/linux/crypto.so)
#add_custom_command( OUTPUT ${legacy_link}
#                    COMMAND ln -s ${legacy_target} ${legacy_link}
#                    DEPENDS install ${legacy_target} 
#                    COMMENT "Generating libcrypto.so symbolic link")

#add_custom_target( install_legacy DEPENDS ${legacy_link} )

#message( STATUS "legacy_link : ${legacy_link}")

#target_link_libraries(common "crypto")
target_link_libraries(libra Threads::Threads common ${CMAKE_CURRENT_LIST_DIR}/libcrypto.so)
#configure_file(${CMAKE_CURRENT_LIST_DIR}/libra ${CMAKE_CURRENT_LIST_DIR}/../test COPYONLY)
#install(TARGETS libra install_legacy DESTINATION bin)
