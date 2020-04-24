
set (LIBNAME "libCollProto")

set (HEADER_FILES
    item.pb.h
    collection.pb.h
)

set (SOURCE_FILES
    item.pb.cc
    collection.pb.cc
)

add_library(${LIBNAME}
    "${HEADER_FILES}"
    "${SOURCE_FILES}"
)

find_package(gRPC CONFIG REQUIRED)
message(STATUS "Using gRPC ${gRPC_VERSION}")

set(_GRPC_GRPCPP gRPC::grpc++)

target_link_libraries( ${LIBNAME}
    ${_GRPC_GRPCPP}
)

target_include_directories(${LIBNAME}
    INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}
)

# install directives
install(TARGETS ${LIBNAME} DESTINATION lib)
install(FILES ${HEADER_FILES} DESTINATION include)