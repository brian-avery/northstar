local nsStream = require("nsStream")

local kafkaSource = {
    Topic = "nsstream-test",
    Brokers = "10.37.13.6:9092",
    ZK = "10.44.6.3:2181"
}

function main()
    local stream, err = nsStream.create("kafka", "description", kafkaSource)
    if err ~= nil then
        error(err)
    end

    err = stream:filter(do_filter_things):map(do_map_things):start()
    if err ~= nil then
        error(err)
    end
end

function do_filter_things(data)
    if data[1] == 79 or data[1] == 111 then
        return true
    end
    return false
end

function do_map_things(data)
    for i=1, #data, 1 do
        if data[i] > 64 and data[i] < 91 then
            data[i] = data[i] + 32
        elseif data[i] > 96 and data[i] < 123 then
            data[i] = data[i] - 32
        end
    end
    return data
end