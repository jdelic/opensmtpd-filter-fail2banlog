#!/bin/bash
go get
go build

fpm \
    -s dir \
    -t deb \
    -p opensmtpd-filter-fail2banlog_0.1.2.deb \
    -n opensmtpd-filter-fail2banlog \
    -v "0.1.2-0" \
    -m "Jonas Maurus" \
    -d "opensmtpd (>=6.8.0)" \
    -d "opensmtpd (<<8.0)" \
    --deb-recommends "fail2ban" \
    --description "Provides logs that can be parsed by fail2ban" \
    --url "https://github.com/jdelic/opensmtpd-filter-fail2banlog" \
    opensmtpd-filter-fail2banlog=/usr/lib/x86_64-linux-gnu/opensmtpd/filter-fail2banlog
