// auto generated 2020-04-15 15:23:32.88375 +0200 CEST m=+1.016444953

package api

import (
	"encoding/base64"
)

var schemas map[string][]byte

func init() {
	schemas = make(map[string][]byte)
	schemas["io.nats.jetstream.api.v1.consumer_configuration"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL2NvbnN1bWVyX2NvbmZpZ3VyYXRpb24uanNvbiIsCiAgImRlc2NyaXB0aW9uIjogIlRoZSBkYXRhIHN0cnVjdHVyZSB0aGF0IGRlc2NyaWJlIHRoZSBjb25maWd1cmF0aW9uIG9mIGEgTkFUUyBKZXRTdHJlYW0gQ29uc3VtZXIiLAogICJ0aXRsZSI6ICJpby5uYXRzLmpldHN0cmVhbS5hcGkudjEuY29uc3VtZXJfY29uZmlndXJhdGlvbiIsCiAgInR5cGUiOiJvYmplY3QiLAogICJkZWZpbml0aW9ucyI6IHsKICAgICJkZWxpdmVyX3BvbGljeSI6IHsKICAgICAgIm9uZU9mIjogWwogICAgICAgIHsiJHJlZiI6ICIjL2RlZmluaXRpb25zL2FsbF9kZWxpdmVyX3BvbGljeSJ9LAogICAgICAgIHsiJHJlZiI6ICIjL2RlZmluaXRpb25zL2xhc3RfZGVsaXZlcl9wb2xpY3kifSwKICAgICAgICB7IiRyZWYiOiAiIy9kZWZpbml0aW9ucy9uZXdfZGVsaXZlcl9wb2xpY3kifSwKICAgICAgICB7IiRyZWYiOiAiIy9kZWZpbml0aW9ucy9zdGFydF9zZXF1ZW5jZV9kZWxpdmVyX3BvbGljeSJ9LAogICAgICAgIHsiJHJlZiI6ICIjL2RlZmluaXRpb25zL3N0YXJ0X3RpbWVfZGVsaXZlcl9wb2xpY3kifQogICAgICBdCiAgICB9LAogICAgImFsbF9kZWxpdmVyX3BvbGljeSI6IHsKICAgICAgInJlcXVpcmVkIjogWyJkZWxpdmVyX3BvbGljeSJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImFsbCJdCiAgICAgICAgfQogICAgICB9CiAgICB9LAogICAgImxhc3RfZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICJyZXF1aXJlZCI6IFsiZGVsaXZlcl9wb2xpY3kiXSwKICAgICAgInByb3BlcnRpZXMiOiB7CiAgICAgICAgImRlbGl2ZXJfcG9saWN5IjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJlbnVtIjogWyJsYXN0Il0KICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAibmV3X2RlbGl2ZXJfcG9saWN5IjogewogICAgICAicmVxdWlyZWQiOiBbImRlbGl2ZXJfcG9saWN5Il0sCiAgICAgICJwcm9wZXJ0aWVzIjogewogICAgICAgICJkZWxpdmVyX3BvbGljeSI6IHsKICAgICAgICAgICJ0eXBlIjogInN0cmluZyIsCiAgICAgICAgICAiZW51bSI6IFsibmV3Il0KICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAic3RhcnRfc2VxdWVuY2VfZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICJyZXF1aXJlZCI6IFsiZGVsaXZlcl9wb2xpY3kiLCAib3B0X3N0YXJ0X3NlcSJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImJ5X3N0YXJ0X3NlcXVlbmNlIl0KICAgICAgICB9LAogICAgICAgICJvcHRfc3RhcnRfc2VxIjogewogICAgICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICAgICAibWluaW11bSI6IDAKICAgICAgICB9CiAgICAgIH0KICAgIH0sCiAgICAic3RhcnRfdGltZV9kZWxpdmVyX3BvbGljeSI6IHsKICAgICAgInJlcXVpcmVkIjogWyJkZWxpdmVyX3BvbGljeSIsICJvcHRfc3RhcnRfdGltZSJdLAogICAgICAicHJvcGVydGllcyI6IHsKICAgICAgICAiZGVsaXZlcl9wb2xpY3kiOiB7CiAgICAgICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAgICAgImVudW0iOiBbImJ5X3N0YXJ0X3RpbWUiXQogICAgICAgIH0sCiAgICAgICAgIm9wdF9zdGFydF90aW1lIjogewogICAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAgICJmb3JtYXQiOiAiZGF0ZS10aW1lIgogICAgICAgIH0KICAgICAgfQogICAgfQogIH0sCiAgInJlcXVpcmVkIjpbCiAgICAiZGVsaXZlcl9wb2xpY3kiLAogICAgImFja19wb2xpY3kiLAogICAgInJlcGxheV9wb2xpY3kiCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgImFsbE9mIjogW3siJHJlZiI6ICIjL2RlZmluaXRpb25zL2RlbGl2ZXJfcG9saWN5In1dLAogICJwcm9wZXJ0aWVzIjogewogICAgImR1cmFibGVfbmFtZSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgInBhdHRlcm4iOiAiXlteLio+XSskIiwKICAgICAgIm1pbkxlbmd0aCI6IDEKICAgIH0sCiAgICAiZGVsaXZlcl9zdWJqZWN0IjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAibWluTGVuZ3RoIjogMQogICAgfSwKICAgICJhY2tfcG9saWN5IjogewogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZW51bSI6IFsibm9uZSIsICJhbGwiLCAiZXhwbGljaXQiXQogICAgfSwKICAgICJhY2tfd2FpdCI6IHsKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJtaW5pbXVtIjogMQogICAgfSwKICAgICJtYXhfZGVsaXZlciI6IHsKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJtaW5pbXVtIjogLTEKICAgIH0sCiAgICAiZmlsdGVyX3N1YmplY3QiOiB7CiAgICAgICJ0eXBlIjogInN0cmluZyIKICAgIH0sCiAgICAicmVwbGF5X3BvbGljeSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImVudW0iOiBbImluc3RhbnQiLCAib3JpZ2luYWwiXQogICAgfSwKICAgICJzYW1wbGVfZnJlcSI6IHsKICAgICAgInR5cGUiOiAic3RyaW5nIgogICAgfQogIH0KfQo=")
	schemas["io.nats.jetstream.api.v1.stream_configuration"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL3N0cmVhbV9jb25maWd1cmF0aW9uLmpzb24iLAogICJkZXNjcmlwdGlvbiI6ICJUaGUgZGF0YSBzdHJ1Y3R1cmUgdGhhdCBkZXNjcmliZSB0aGUgY29uZmlndXJhdGlvbiBvZiBhIE5BVFMgSmV0U3RyZWFtIFN0cmVhbSIsCiAgInRpdGxlIjogImlvLm5hdHMuamV0c3RyZWFtLmFwaS52MS5zdHJlYW1fY29uZmlndXJhdGlvbiIsCiAgInR5cGUiOiJvYmplY3QiLAogICJyZXF1aXJlZCI6WwogICAgIm5hbWUiLAogICAgInJldGVudGlvbiIsCiAgICAibWF4X2NvbnN1bWVycyIsCiAgICAibWF4X21zZ3MiLAogICAgIm1heF9ieXRlcyIsCiAgICAibWF4X2FnZSIsCiAgICAic3RvcmFnZSIsCiAgICAibnVtX3JlcGxpY2FzIgogIF0sCiAgImFkZGl0aW9uYWxJdGVtcyI6IGZhbHNlLAogICJwcm9wZXJ0aWVzIjogewogICAgIm5hbWUiOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJBIHVuaXF1ZSBuYW1lIGZvciB0aGUgU3RyZWFtLiIsCiAgICAgICJtaW5MZW5ndGgiOiAxLAogICAgICAicGF0dGVybiI6ICJeW14uKj5dKyQiCiAgICB9LAogICAgInN1YmplY3RzIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiQSBsaXN0IG9mIHN1YmplY3RzIHRvIGNvbnN1bWUsIHN1cHBvcnRzIHdpbGRjYXJkcy4iLAogICAgICAidHlwZSI6ICJhcnJheSIsCiAgICAgICJtaW5MZW5ndGgiOiAxLAogICAgICAiaXRlbXMiOiB7CiAgICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgICAibWluTGVuZ3RoIjogMQogICAgICB9CiAgICB9LAogICAgInJldGVudGlvbiI6IHsKICAgICAgImRlc2NyaXB0aW9uIjogIkhvdyBtZXNzYWdlcyBhcmUgcmV0YWluZWQgaW4gdGhlIFN0cmVhbSwgb25jZSB0aGlzIGlzIGV4Y2VlZGVkIG9sZCBtZXNzYWdlcyBhcmUgcmVtb3ZlZC4iLAogICAgICAidHlwZSI6ICJzdHJpbmciLAogICAgICAiZW51bSI6IFsibGltaXRzIiwgImludGVyZXN0IiwgIndvcmtxdWV1ZSJdLAogICAgICAiZGVmYXVsdCI6ICJsaW1pdHMiCiAgICB9LAogICAgIm1heF9jb25zdW1lcnMiOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJIb3cgbWFueSBDb25zdW1lcnMgY2FuIGJlIGRlZmluZWQgZm9yIGEgZ2l2ZW4gU3RyZWFtLiAtMSBmb3IgdW5saW1pdGVkLiIsCiAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAibWluaW11bSI6IC0xLAogICAgICAiZGVmYXVsdCI6IC0xCiAgICB9LAogICAgIm1heF9tc2dzIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiSG93IG1hbnkgbWVzc2FnZXMgbWF5IGJlIGluIGEgU3RyZWFtLCBvbGRlc3QgbWVzc2FnZXMgd2lsbCBiZSByZW1vdmVkIGlmIHRoZSBTdHJlYW0gZXhjZWVkcyB0aGlzIHNpemUuIC0xIGZvciB1bmxpbWl0ZWQuIiwKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJtaW5pbXVtIjogLTEsCiAgICAgICJkZWZhdWx0IjogLTEKICAgIH0sCiAgICAibWF4X2J5dGVzIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiSG93IGJpZyB0aGUgU3RyZWFtIG1heSBiZSwgd2hlbiB0aGUgY29tYmluZWQgc3RyZWFtIHNpemUgZXhjZWVkcyB0aGlzIG9sZCBtZXNzYWdlcyBhcmUgcmVtb3ZlZC4gLTEgZm9yIHVubGltaXRlZC4iLAogICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgIm1pbmltdW0iOiAtMSwKICAgICAgImRlZmF1bHQiOiAtMQogICAgfSwKICAgICJtYXhfYWdlIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiTWF4aW11bSBhZ2Ugb2YgYW55IG1lc3NhZ2UgaW4gdGhlIHN0cmVhbSwgZXhwcmVzc2VkIGluIG1pY3Jvc2Vjb25kcy4gLTEgZm9yIHVubGltaXRlZC4iLAogICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgIm1pbmltdW0iOiAwLAogICAgICAiZGVmYXVsdCI6IDAKICAgIH0sCiAgICAibWF4X21zZ19zaXplIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIGxhcmdlc3QgbWVzc2FnZSB0aGF0IHdpbGwgYmUgYWNjZXB0ZWQgYnkgdGhlIFN0cmVhbS4gLTEgZm9yIHVubGltaXRlZC4iLAogICAgICAidHlwZSI6ICJpbnRlZ2VyIiwKICAgICAgIm1pbmltdW0iOiAtMSwKICAgICAgImRlZmF1bHQiOiAtMQogICAgfSwKICAgICJzdG9yYWdlIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiVGhlIHN0b3JhZ2UgYmFja2VuZCB0byB1c2UgZm9yIHRoZSBTdHJlYW0uIiwKICAgICAgInR5cGUiOiAic3RyaW5nIiwKICAgICAgImVudW0iOiBbImZpbGUiLCAibWVtb3J5Il0KICAgIH0sCiAgICAibnVtX3JlcGxpY2FzIjogewogICAgICAiZGVzY3JpcHRpb24iOiAiSG93IG1hbnkgcmVwbGljYXMgdG8ga2VlcCBmb3IgZWFjaCBtZXNzYWdlLiIsCiAgICAgICJ0eXBlIjogImludGVnZXIiLAogICAgICAibWluaW11bSI6IDEsCiAgICAgICJkZWZhdWx0IjogMQogICAgfSwKICAgICJub19hY2siOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJEaXNhYmxlcyBhY2tub3dsZWRnaW5nIG1lc3NhZ2VzIHRoYXQgYXJlIHJlY2VpdmVkIGJ5IHRoZSBTdHJlYW0uIiwKICAgICAgInR5cGUiOiAiYm9vbGVhbiIsCiAgICAgICJkZWZhdWx0IjogZmFsc2UKICAgIH0sCiAgICAidGVtcGxhdGVfb3duZXIiOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJXaGVuIHRoZSBTdHJlYW0gaXMgbWFuYWdlZCBieSBhIFN0cmVhbSBUZW1wbGF0ZSB0aGlzIGlkZW50aWZpZXMgdGhlIHRlbXBsYXRlIHRoYXQgbWFuYWdlcyB0aGUgU3RyZWFtLiIsCiAgICAgICJ0eXBlIjogInN0cmluZyIKICAgIH0KICB9Cn0K")
	schemas["io.nats.jetstream.api.v1.stream_template_configuration"], _ = base64.StdEncoding.DecodeString("ewogICIkc2NoZW1hIjogImh0dHA6Ly9qc29uLXNjaGVtYS5vcmcvZHJhZnQtMDcvc2NoZW1hIyIsCiAgIiRpZCI6ICJodHRwczovL25hdHMuaW8vc2NoZW1hcy9qZXRzdHJlYW0vYXBpL3YxL3N0cmVhbV90ZW1wbGF0ZV9jb25maWd1cmF0aW9uLmpzb24iLAogICJkZXNjcmlwdGlvbiI6ICJUaGUgZGF0YSBzdHJ1Y3R1cmUgdGhhdCBkZXNjcmliZSB0aGUgY29uZmlndXJhdGlvbiBvZiBhIE5BVFMgSmV0U3RyZWFtIFN0cmVhbSBUZW1wbGF0ZSIsCiAgInRpdGxlIjogImlvLm5hdHMuamV0c3RyZWFtLmFwaS52MS5zdHJlYW1fdGVtcGxhdGVfY29uZmlndXJhdGlvbiIsCiAgInR5cGUiOiJvYmplY3QiLAogICJyZXF1aXJlZCI6WwogICAgIm5hbWUiLAogICAgImNvbmZpZyIsCiAgICAibWF4X3N0cmVhbXMiCiAgXSwKICAiYWRkaXRpb25hbEl0ZW1zIjogZmFsc2UsCiAgInByb3BlcnRpZXMiOiB7CiAgICAibmFtZSI6IHsKICAgICAgImRlc2NyaXB0aW9uIjogIkEgdW5pcXVlIG5hbWUgZm9yIHRoZSBTdHJlYW0gVGVtcGxhdGUuIiwKICAgICAgIm1pbkxlbmd0aCI6IDEsCiAgICAgICJwYXR0ZXJuIjogIl5bXi4qPl0rJCIKICAgIH0sCiAgICAibWF4X3N0cmVhbXMiOiB7CiAgICAgICJkZXNjcmlwdGlvbiI6ICJUaGUgbWF4aW11bSBudW1iZXIgb2YgU3RyZWFtcyB0aGlzIFRlbXBsYXRlIGNhbiBjcmVhdGUsIC0xIGZvciB1bmxpbWl0ZWQuIiwKICAgICAgInR5cGUiOiAiaW50ZWdlciIsCiAgICAgICJtaW5pbXVtIjogLTEsCiAgICAgICJkZWZhdWx0IjogLTEKICAgIH0sCiAgICAiY29uZmlnIjogewogICAgICAiJHJlZiI6ICJzdHJlYW1fY29uZmlndXJhdGlvbi5qc29uIgogICAgfQogIH0KfQo=")
}
