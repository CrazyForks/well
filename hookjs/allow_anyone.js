define(["collections/peers"], () => {
  return {
    onRecordCreateRequest(e) {
      let r = e.record
      if (
        r.getString("name") == "" &&
        r.getBool("disabled") == true &&
        r.getString("ipv6") == ""
      ) {
        r.set("disabled", false)
        r.set("ipv6", "auto")
      }
      return e.next()
    },
  }
})
