A filter to write fail2ban compatible logs for OpenSMTPD
========================================================

OpenSMTPD logs the caller's ip address and authentication results on separate
lines in the log file. This filter creates a combined log entry that can be
matched by fail2ban.

How to use this
---------------

Once the packages are uploaded, you will be able to install on Debian
like this:

::

    wget -O /etc/apt/kryrings/maurusnet-archive-keyring.gpg http://repo.maurus.net/02CBD940A78049AF.pem
    echo "deb [signed-by=/etc/apt/keyrings/maurusnet-archive-keyring.gpg] http://repo.maurus.net/release/trixie/ mn-release main" > /etc/apt/sources.list.d/maurusnet.list
    apt update
    apt install opensmtpd-filter-fail2banlog


Example usage in smtpd.conf
---------------------------

In your OpenSMTPD configuration activate ``filter-greylistd``:

::

    filter "fail2banlog" proc-exec "/usr/lib/x86_64-linux-gnu/opensmtpd/filter-fail2banlog"
    listen on "127.0.0.1" port 25 filter fail2banlog


Example usage for fail2ban
--------------------------

``/etc/fail2ban/filter.d/opensmtpd-auth.conf``:

::

    [Definition]
    failregex = ^.*opensmtpd-f2b: auth-failure rip=<HOST> user=".*" rdns=".*" result=".*" session=.*$
    ignoreregex =
    journalmatch = _SYSTEMD_UNIT=opensmtpd.service


``/etc/fail2ban/jail.d/opensmtpd.local``:

::

    [opensmtpd-auth]
    enabled = true
    filter = opensmtpd-auth
    backend = systemd
    journalmatch = _SYSTEMD_UNIT=opensmtpd.service
    port = 25,465,587
    findtime = 10m
    maxretry = 5
    bantime = 1h
    banaction = nftables


.. _osfgo: https://github.com/jdelic/opensmtpd-filters-go
