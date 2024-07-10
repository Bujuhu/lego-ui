This was a university project.


# lego-ui

lego-ui is a graphical user interface written in go developed on top of the [go-acme/lego](https://github.com/go-acme/lego) library to make it simple for everyone to get 
Lets Encrypt certificates using the dns-01 acme challenge

## Vision

For novices, it can be pretty difficult to generate new TLS certificates for their devices. While there are some great solutions like from [Synology](https://kb.synology.com/de-de/DSM/tutorial/How_to_enable_HTTPS_and_create_a_certificate_signing_request_on_your_Synology_NAS), [Cpanel](https://blog.cpanel.com/how-to-configure-and-manage-lets-encrypt-in-cpanel/), [Nginx Proxy Manager](https://nginxproxymanager.com/) or even [GitHub](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https) itself, which are great to use but force the user to host their content in certain (sometimes proprietary) environments or using specific software.

If users want to create TLS certificates, they need to make a ton of network precautions, like opening a port (which some isps won't even allow), or configuring a reverse web proxy (which requires an additional provider or cloud server). 

To improve this process, we want to create a User Interface, using the dns-acme challenge. The downside of this is, that the user needs a domain. Only authentication data and personal information to generate the TLS Certificate are needed for this process.

### Mockup

![Mockup of the Graphical User Interface](/docs/p1/mockup.jpg)

### Screenshots

- 2022-12-20

![](/docs/p1/screenshot-2022-12-20.png)

## Implementation
We want to implement the application using the go programming language, using the [go-acme/lego](https://github.com/go-acme/lego) library and
[fyne-io/fyne](https://github.com/fyne-io/fyne) as our gui library.
Lego provides a go library to make DNS challanges out of the box to loads of different DNS providers.
We decided on fyne as our GUI library, because it
- Runs natively on the machine (which means better integration into system features like keyring and the file system)
- Well-defined design template (reducing amount of work needed to design the UX)


We want to implement the following tasks during this project:
- Save Personal Information to use for certificates
- Save DNS-API credentials using the keyring (if keyring exists)
- Put the certificate into the local keyring or export it to the local file system if no keyring exists
- Certificate Expiration Warning
- Renew Certificate
- Export Certifcate

### Changes
We've previously stated, that we want to make it possible to automatically configure the host file using elektra based on the domain name of a generated certificate. This however would require elevated privileges for automated setup, so we will this further back to explore, how/if we can implement this feature safely

The way that multiple DNS providers are implemented within go-acme/lego, makes it difficult to implement each individual DNS provider. As there are so many, we will probably only implement a handful of DNS providers.

## Developer Documentation & Reducing Entry Barriers
We will 

## User Documentation
