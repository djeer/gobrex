# gobrex
bittrex.com command-line currency monitoring
Приложение получает по API Bittrex (https://bittrex.com/home/api) тикер /public/getticker по парам BTC-ETH, BTC-LTC, BTC-XMR, BTC-NXT, BTC-DASH (пары должны легко добавляться и удаляться в коде или в конфигурационном файле) и, если изменилось значение Last, выводить его в лог. Запросы должны выполняться параллельно, но с ограничением 3 запроса в секунду (это значение должно легко меняться на любое другое).
