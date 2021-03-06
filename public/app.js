var new_short = new Vue({
    el: '#new_short',
    data: {
        requestUrl: document.location.href,
        message: 'get short url !',
        originalUrl: '',
        userShort: '',
        haveShort: false,
        yourShortUrl: '',
        saySomething: '',
        getSuccess: false,
    },
    methods: {
        createUrl: async function () {
            // send request 
            const data = `shortname=${this.userShort}&&url=${this.originalUrl}`
            const response = await fetch(this.requestUrl+'api', {
                method: 'post',
                headers: {
                  "Content-type": "application/x-www-form-urlencoded; charset=UTF-8"
                },
                body: data
            });

            const result = await response.json()
            this.haveShort = true

            if (result.checkCode == 100) {
                // fail
                this.getSuccess = false
                this.userShort = result.short_name
                this.saySomething = `Your desired short url had been occupied !\n Maybe use  ${result.short_name}  ?
                    Press button to create short url !`
                this.yourShortUrl = ''
            } else {
                // work !
                this.getSuccess = true
                this.originalUrl = ''
                this.userShort = ''
                this.saySomething = "\n\n\n\nThis is your url :"
                this.yourShortUrl = `${this.requestUrl}${result.short_name}`
            }
        }
    }
})