local nsQL = require("nsQL")

function main()
    local query = [[
        UPDATE  nssim.sampleData
        SET     sampleData.createdtime = '2017-01-01 10:10:10'
        WHERE   sampleData.rowId = 'aca7ae94-1fc9-11e7-93ae-92361f001953'
                AND sampleData.id = '9e3a6e50-1fc9-11e7-93ae-92361f001953';
    ]]
    local source = {
        Protocol = "cassandra",
        Host = "10.32.49.6",
        Port = "31838",
        Backend = "native"
    }
    local options = {}
    local connection = createConnection(source)
    processQuery(connection, query, options)
    teardownConnection(connection)
end

function createConnection(source)
    local connection, err = nsQL.connect(source)
    if(err ~= nil) then
        error(err)
    end
    return connection
end

function teardownConnection(connection)
    local err = connection:disconnect()
    if(err ~= nil) then
        error(err)
    end
end

function processQuery(connection, query, options)
    local resp, err = connection:query(query, options)
    if(err ~= nil) then
        error(err)
    end
    return resp
end