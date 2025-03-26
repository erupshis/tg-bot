// Code generated by locgen. DO NOT EDIT.

package locales

type MessageKey string

var (
	Errors = struct {
	}{}

	Messages = struct {
		Admin struct {
			NewMessage struct {
				MessageHeader MessageKey
				ApproveButton MessageKey
				RejectButton  MessageKey
			}
			Callback struct {
				Approved MessageKey
				Rejected MessageKey
			}
		}
		User struct {
			MessageTooShort MessageKey
			MessageReceived MessageKey
		}
		Commands struct {
			Help  MessageKey
			Start MessageKey
		}
	}{
		Admin: struct {
			NewMessage struct {
				MessageHeader MessageKey
				ApproveButton MessageKey
				RejectButton  MessageKey
			}
			Callback struct {
				Approved MessageKey
				Rejected MessageKey
			}
		}{
			NewMessage: struct {
				MessageHeader MessageKey
				ApproveButton MessageKey
				RejectButton  MessageKey
			}{
				MessageHeader: "messages.admin.new_message.message_header",
				ApproveButton: "messages.admin.new_message.approve_button",
				RejectButton:  "messages.admin.new_message.reject_button",
			},
			Callback: struct {
				Approved MessageKey
				Rejected MessageKey
			}{
				Approved: "messages.admin.callback.approved",
				Rejected: "messages.admin.callback.rejected",
			},
		},
		User: struct {
			MessageTooShort MessageKey
			MessageReceived MessageKey
		}{
			MessageTooShort: "messages.user.message_too_short",
			MessageReceived: "messages.user.message_received",
		},
		Commands: struct {
			Help  MessageKey
			Start MessageKey
		}{
			Help:  "messages.commands.help",
			Start: "messages.commands.start",
		},
	}
)
