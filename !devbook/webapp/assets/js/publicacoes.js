$("#nova-publicacao").on("submit", criarPublicacao)

// Importante fazer quando a classe pode mudar
$(document).on("click", ".curtir-publicacao", curtirPublicacao)
$(document).on("click", ".descurtir-publicacao", descurtirPublicacao)
$("#atualizar-publicacao").on("click", atualizarPublicacao)
$(".deletar-publicacao").on("click", deletarPublicacao)

function criarPublicacao(evento) {
    evento.preventDefault()

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $("#titulo").val(),
            conteudo: $("#conteudo").val(),
        }
    }).done(function () {
        window.location = "/home"
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao criar publicação", "error")
    })
}

function curtirPublicacao(evento) {
    const elementoClicado = $(evento.target)
    // Sobe no elemento clicado até encontrar um div, então pega a data
    const publicacaoID = elementoClicado.closest("div").data("publicacao-id")

    elementoClicado.prop("disabled", true)

    $.ajax({
        url: `/publicacoes/${publicacaoID}/curtir`,
        method: "POST"
    }).done(function () {
        const contadorDeCurtidas = elementoClicado.next("span")
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1)

        elementoClicado.addClass("descurtir-publicacao")
        elementoClicado.addClass("text-danger")
        elementoClicado.removeClass("curtir-publicacao")
    }).fail(function () {
        Swal.fire("Ops...", "Houve um erro na conexão, tente novamente", "error")
    }).always(function () {
        elementoClicado.prop("disabled", false)
    })
}

function descurtirPublicacao(evento) {
    const elementoClicado = $(evento.target)
    // Sobe no elemento clicado até encontrar um div, então pega a data
    const publicacaoID = elementoClicado.closest("div").data("publicacao-id")

    elementoClicado.prop("disabled", true)

    $.ajax({
        url: `/publicacoes/${publicacaoID}/descurtir`,
        method: "POST"
    }).done(function () {
        const contadorDeCurtidas = elementoClicado.next("span")
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1)

        elementoClicado.addClass("curtir-publicacao")
        elementoClicado.removeClass("text-danger")
        elementoClicado.removeClass("descurtir-publicacao")
    }).fail(function () {
        Swal.fire("Ops...", "Houve um erro na conexão, tente novamente", "error")
    }).always(function () {
        elementoClicado.prop("disabled", false)
    })
}

function atualizarPublicacao(evento) {
    $(this).prop("disabled", true)

    const publicacaoId = $(this).data("publicacao-id")

    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $("#titulo").val(),
            conteudo: $("#conteudo").val(),
        }
    }).done(function () {
        Swal.fire("Sucesso", "Publicação Editada com sucesso!", "success")
            .then(function () {
                window.location = "/home"
            })
    }).fail(function () {
        Swal.fire("Ops...", "Erro ao atualizar publicação", "error")
    }).always(function () {
        $("#atualizar-publicacao").prop("disabled", false)
    })
}

function deletarPublicacao(evento) {
    evento.preventDefault()

    Swal.fire({
        title: "Atenção!",
        text: "Deseja mesmo deletar esta publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function (confirmacao) {
        if (!confirmacao.value) {
            return
        } else {
            const elementoClicado = $(evento.target)
            // Sobe no elemento clicado até encontrar um div, então pega a data
            const publicacao = elementoClicado.closest("div")
            const publicacaoID = publicacao.data("publicacao-id")

            elementoClicado.prop("disabled", true)

            $.ajax({
                url: `/publicacoes/${publicacaoID}`,
                method: "DELETE"
            }).done(function () {
                publicacao.fadeOut("slow", function () {
                    $(this).remove()
                })
            }).fail(function () {
                Swal.fire("Ops...", "Erro ao excluir publicação", "error")
            })
        }
    })
}