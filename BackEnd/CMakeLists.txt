cmake_minimum_required (VERSION 3.2)

#add lib subdirectory to build
add_subdirectory (src)

project (embedded_server Go)
set(CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}/cmake")


#include directories to building path
include_directories(${PROJECT_SOURCE_DIR}/src/include)
include_directories(${PROJECT_SOURCE_DIR}/src/lib)
include_directories(${PROJECT_SOURCE_DIR}/route)
include_directories(${PROJECT_SOURCE_DIR}/controllers)
include_directories(${PROJECT_SOURCE_DIR}/resources)

set(SOURCE_FILES main.go)

add_custom_command(
    OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/main.go
    COMMAND go build ${SOURCE_FILES}
)

add_executable(home_server main.go)

# target_link_libraries(pinControl)
#install the binary
# install (TARGETS pinControl DESTINATION bin)
#package source generation
set(MAJOR "0")
set(MINOR "1")
set(PATCH "1")
set(CPACK_SOURCE_GENERATOR "TGZ")
set(CPACK_SOURCE_PACKAGE_FILE_NAME "${CMAKE_PROJECT_NAME}-${MAJOR}.${MINOR}.${PATCH}")
set(CPACK_SOURCE_IGNORE_FILES "/build/;${CPACK_SOURCE_IGNORE_FILES}")
include(CPack)
