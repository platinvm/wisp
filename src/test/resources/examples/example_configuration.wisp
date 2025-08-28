/*
    this is an example of what it should
    look like to write a configuraton
*/

{
    device: {
        id: 0xA5,
        type: "temperature_sensor",
        mac: b8:27:eb:01:23:45,
        firmware: 1.0.3,
        license: true,
        network: {
            ipv4: 10.160.6.119,
            ipv6: fe80::5008:bf87:9339:47ad,
        },
        settings: {
            timeout: 10m,
            sampling_interval: 2.5s,
            threshold: 25.0,
            status_led_colours: {
                normal: #00ff00,
                warning: #ffff00,
                critical: #ff0000
            },
            calibration: {
                offset: 0b00010101, // 21
                factor: 1.075,
            }
        }
    }
}