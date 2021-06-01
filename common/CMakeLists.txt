cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
        
set(CRYPTO_DIR
        $ENV{COMMON_ROOT}/lib/linux)

        
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT} ${LIBRA_HDR} ${CRYPTO_DIR})

#find_library(crypto_lib
#         NAMES 
#            crypto
#         PATHS 
#            ${CRYPTO_DIR}
#)

#add_library(crypto SHARED IMPORTED)
#if ( NOT CRYPTO_LIB )
#        message( WARNING "CRYPTO_LIB NOT FOUND in common")
#else()
#        message( STATUS "CRYPTO_LIB : ${CRYPTO_LIB}")
#endif( NOT CRYPTO_LIB )

target_link_libraries(common crypto)

install(TARGETS common DESTINATION lib)
