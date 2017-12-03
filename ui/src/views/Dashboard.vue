<template lang="pug">
  div.animated.fadeIn
    b-row
        b-col(md='12')
            b-card(header='BGP Sessions')
              b-row
                b-col(sm='12')
                  b-row
                    b-col(sm='3')
                        small.text-muted Total Sessions
                        br
                        |
                        strong.h4 {{ sessions.length }}
                        |
                        .chart-wrapper
                    b-col(sm='3')
                        small.text-muted Filtered
                        br  
                        |
                        strong.h4 {{ filteredSessions.length }}
                        |
                        .chart-wrapper 
                    b-col(sm='3')
                        small.text-muted Established
                        br
                        |
                        strong.h4 {{ upSessions.length }}
                        |
                        .chart-wrapper
                    b-col(sm='3')                  
                        small.text-muted Down
                        br
                        |                     
                        strong.h4 {{ downSessions.length }}
                        |                     
                        .chart-wrapper
            b-card(header='BGP Sessions')
                b-form-fieldset()
                  b-form-input(type="text" v-model="searchQuery" id="name" placeholder="Search")
                b-table(:items="filteredSessions" :fields="fields" hover)
                  div(slot="host" slot-scope="item")
                    div {{item.value}} 
                    div.small.text-muted VRF: {{ item.item['peer-cfg-rti'] }}
                    div.small.text-muted Group: {{ item.item['peer-group'] }}
                  div(slot="local-address" slot-scope="item")
                    div {{item.value}} 
                    div.small.text-muted AS: {{ item.item['local-as'] }}
                    div.small.text-muted Type: {{ item.item['peer-type'] }}    
                  div(slot="peer-address" slot-scope="item")
                    div {{item.value}}
                    div.small.text-muted AS: {{ item.item['peer-as'] }}
                    div.small.text-muted Flag: {{ item.item['peer-flags'] }}
                  div(slot="peer-state" slot-scope="item")
                    b-badge(:variant="getStateBadge(item.value)") {{item.value}}
                    div.small.text-muted Last err: {{ item.item['last-error'] }}    
                    div.small.text-muted Last event: {{ item.item['last-event'] }}      
               
</template>

<script>
import axios from 'axios'

export default {
  name: 'dashboard',
  data () {
    return {
      searchQuery: '',
      sessions: [],
      errors: [],
      fields: ['host', 'local-address', 'peer-address', 'peer-state', 'err']
    }
  },
  // Fetches posts when the component is created.
  created () {
    axios
      .get('http://' + window.location.hostname + ':8888/bgpnei')
      .then(response => {
        // JSON responses are automatically parsed.
        var sessions = response.data
        for (var hidx = sessions.length - 1; hidx >= 0; hidx--) {
          if (sessions[hidx]['bgp-peer'] !== null) {
            for (var nidx = sessions[hidx]['bgp-peer'].length - 1; nidx >= 0; nidx--) {
              var nei = sessions[hidx]['bgp-peer'][nidx]
              nei.host = sessions[hidx].host
              this.sessions.push(nei)
            }
          } else if (sessions[hidx].err !== '') {
            this.sessions.push(sessions[hidx])
          }
        }
      })
      .catch(err => {
        this.errors.push(err)
        console.log(err)
      })
  },
  methods: {
    getStateBadge (status) {
      return status === 'Connect' ? 'warning'
        : status === 'Inactive' ? 'secondary'
          : status === 'Active' ? 'warning'
            : status === 'Idle' ? 'danger' : 'primary'
    }
  },
  computed: {
    filteredSessions: function () {
      var self = this
      return self.sessions.filter(function (session) {
        var searchRegex = new RegExp(self.searchQuery, 'i')
        return searchRegex.test(session.host) || searchRegex.test(session['peer-address'])
      })
    },
    upSessions: function () {
      var self = this
      return self.sessions.filter(function (session) {
        if (session['peer-state'] === 'Established') return session
      })
    },
    downSessions: function () {
      var self = this
      return self.sessions.filter(function (session) {
        if (session['peer-state'] !== 'Established') return session
      })
    }
  }
  // mock () {
  //   var jsonString = ''
  //   var s = jsonString.replace(/\\n/g, '\\n')
  //     .replace(/\\'/g, '\\\'')
  //     .replace(/\\"/g, '\\"')
  //     .replace(/\\&/g, '\\&')
  //     .replace(/\\r/g, '\\r')
  //     .replace(/\\t/g, '\\t')
  //     .replace(/\\b/g, '\\b')
  //     .replace(/\\f/g, '\\f')
  //   s = s.replace(/[\u0000-\u0019]+/g, '')
  //   var sessions = JSON.parse(s)
  //   for (var hidx = sessions.length - 1; hidx >= 0; hidx--) {
  //     if (sessions[hidx]['bgp-peer'] !== null) {
  //       for (var nidx = sessions[hidx]['bgp-peer'].length - 1; nidx >= 0; nidx--) {
  //         var nei = sessions[hidx]['bgp-peer'][nidx]
  //         nei.host = sessions[hidx].host
  //         this.sessions.push(nei)
  //       }
  //     }
  //   }
  // }
}
</script>
