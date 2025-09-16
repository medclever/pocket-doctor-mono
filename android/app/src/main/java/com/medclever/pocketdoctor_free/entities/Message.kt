package com.medclever.pocketdoctor_free.entities

import kotlinx.serialization.Serializable

@Serializable
data class MessageData(
    val messageId: String,
    val message: String,
    val dateCreate: String
)

class Message(val message: String) {
}