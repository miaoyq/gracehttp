from busybox

copy ./http-demo /

cmd ["server"]
entrypoint ["/http-demo"]
